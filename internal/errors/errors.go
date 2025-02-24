// errors/errors.go
package errors

import "fmt"

// NotFoundError 表示资源未找到的错误
type NotFoundError struct {
	Resource string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s 不存在", e.Resource)
}

// NewNotFoundError 创建一个新的 NotFoundError
func NewNotFoundError(resource string) error {
	return &NotFoundError{Resource: resource}
}
