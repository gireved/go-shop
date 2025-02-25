package startup

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	conf "go-shop/config"
	"go-shop/internal/models"
	zaplogger "go-shop/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
	"sync"
	"time"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

func InitMySQL() error {
	var initErr error
	// 1. 配置校验
	dbOnce.Do(func() {
		mConfig, exists := conf.Config.Mysql["default"]
		if !exists {
			initErr = fmt.Errorf("数据库默认连接未找到")
			return
		}

		if err := validateConfig(mConfig); err != nil {
			initErr = fmt.Errorf("无效的mysql配置：%w", err)
		}

		// 2. 构造DSN（安全方式）
		dsn := buildDSN(mConfig)

		// 3. 初始化主连接
		primaryDB, err := createGormDB(dsn, mConfig)
		if err != nil {
			initErr = fmt.Errorf("创建连接失败: %w", err)
			return
		}

		// 4. 配置读写分离（示例）
		if err := setupReadWriteSeparation(primaryDB, mConfig); err != nil {
			initErr = fmt.Errorf("设置读写分离失败: %w", err)
			return
		}

		// 5. 配置连接池
		if err := configureConnectionPool(primaryDB, mConfig); err != nil {
			initErr = fmt.Errorf("配置连接池失败: %w", err)
			return
		}

		// **6. 进行自动迁移**
		if err := autoMigrateTables(primaryDB); err != nil {
			initErr = fmt.Errorf("数据库迁移失败: %w", err)
			return
		}

		dbInstance = primaryDB
	})
	return initErr
}

// GetDB 安全获取数据库实例
func GetDB() *gorm.DB {
	if dbInstance == nil {
		panic("database not initialized")
	}
	return dbInstance
}

// 校验配置完整性
func validateConfig(c *conf.Mysql) error {

	requiredFields := map[string]string{
		"Dialect":  c.Dialect,
		"DBHost":   c.DBHost,
		"DBPort":   c.DBPort,
		"DBName":   c.DBName,
		"UserName": c.UserName,
		"Password": c.Password,
		"Charset":  c.Charset,
	}

	for name, value := range requiredFields {
		if value == "" {
			return fmt.Errorf("没有找到配置%s", name)
		}
	}

	if c.MaxOpenConns < 1 || c.MaxIdleConns < 0 {
		return fmt.Errorf("无效的连接池设置")
	}

	return nil
}

// 安全构建连接字符串
func buildDSN(c *conf.Mysql) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local&timeout=5s",
		c.UserName,
		c.Password,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.Charset)
}

// 创建GORM数据库实例
func createGormDB(dsn string, c *conf.Mysql) (*gorm.DB, error) {
	// 根据环境配置日志器
	gormConfig := &gorm.Config{
		Logger: newCustomLogger(),
		// 命名策略：使用单数表名
		// 例如：User模型对应的表名为"user"而非默认的"users"
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction: true, // 禁用默认事务提升性能
		PrepareStmt:            true, // 开启预变编译
		NowFunc: func() time.Time {
			return time.Now().UTC().Truncate(time.Microsecond) // 对齐MySQL 8.0的时间精度
		},
		DisableAutomaticPing: false, // 自动保活连接
	}

	return gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DriverName:                "mysql", // 明确指定使用8.x驱动
		DefaultStringSize:         191,     // 适配utf8mb4最大索引长度
		DisableDatetimePrecision:  false,   // 启用datetime精度
		DontSupportRenameIndex:    false,   // 支持在线DDL
		DontSupportRenameColumn:   false,   // 支持列重命名
		SkipInitializeWithVersion: true,    // 禁用版本自动检测
		Conn:                      nil,     // 可自定义底层连接
	}), gormConfig)
}

// 配置连接池参数
func configureConnectionPool(db *gorm.DB, c *conf.Mysql) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(c.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(c.ConnMaxIdleTime)

	// 启动连接健康检查
	//go monitorConnectionPool(sqlDB)
	return nil
}

// 配置读写分离
func setupReadWriteSeparation(db *gorm.DB, c *conf.Mysql) error {
	// 从配置中获取读写节点信息
	readDNSs := []string{c.ReadReplica1, c.ReadReplica2}
	var replicas []gorm.Dialector

	for _, dns := range readDNSs {
		if dns != "" {
			replicas = append(replicas, mysql.Open(dns))
		}
	}

	if len(replicas) == 0 {
		return nil
	}

	return db.Use(dbresolver.Register(dbresolver.Config{
		Sources:           []gorm.Dialector{mysql.Open(c.WriteDSN)},
		Replicas:          replicas,
		Policy:            dbresolver.RandomPolicy{},
		TraceResolverMode: true, // 开启解析器追踪
	}))
}

// 自定义日志记录器
func newCustomLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             500 * time.Millisecond,
			Colorful:                  gin.Mode() == gin.DebugMode,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
		},
	)
}

// 自动迁移数据库表
func autoMigrateTables(db *gorm.DB) error {
	zaplogger.Info("开始数据库表迁移...")

	err := db.AutoMigrate(
		&models.Product{}, // 商品表
	)

	if err != nil {
		return err
	}

	zaplogger.Info("数据库表迁移完成")
	return nil
}

//  连接池健康监控
/*func monitorConnectionPool(sqlDB *sql.DB) {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		stats := sqlDB.Stats()
		metrics.ReportConnectionPoolStats(stats)

		if stats.OpenConnections >= stats.MaxOpenConnections-5 {
			alert.Send("DB connection pool reaching capacity")
		}
	}
}*/
