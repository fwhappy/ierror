package ierror

import (
	"fmt"
	"runtime/debug"

	"github.com/fwhappy/iutil"
)

// Error 错误
type Error struct {
	code  int
	msg   string // 带参数的具体错误描述
	error        // 返还给客户端的错误
}

// 错误配置
var cfg map[int][]string

func init() {
	cfg = map[int][]string{}
}

// NewError 新建一个错误对象
// 若错误号未配置，返回一个通用描述
func NewError(code int, argv ...interface{}) *Error {
	err := &Error{code: code}
	if msg, exists := cfg[code]; exists {
		if len(msg) >= 2 {
			err.msg = fmt.Sprintf(msg[1], argv...)
		} else {
			err.msg = fmt.Sprintf(msg[0], argv...)
		}
		err.error = fmt.Errorf(msg[0])
	} else {
		err.msg = fmt.Sprintf("错误号[%v]未定义", code)
		err.error = fmt.Errorf("请求失败")
	}
	return err
}

// GetCode 获取错误号
func (err *Error) GetCode() int {
	return err.code
}

// GetMsg 获取详细的错误描述
func (err *Error) GetMsg() string {
	return err.msg
}

// MustNil 判断error是否是nil
func MustNil(err error) bool {
	if err != nil {
		fmt.Printf("[ %s ]Error: %s\nstack:\n%s\n", iutil.GetTimestamp(), err.Error(), debug.Stack())
		// core.Logger.Error(msg)
		return false
	}
	return true
}

// IMustNil 判断ierror是否为nil
func IMustNil(err *Error) bool {
	if err != nil {
		// msg := fmt.Sprintf("Error: %v, code:%v\nstack:\n%s", err.GetMsg(), err.GetCode(), debug.Stack())
		fmt.Printf("[ %s ]Error: %s\nstack:\n%s\n", iutil.GetTimestamp(), err.Error(), debug.Stack())
		return false
	}
	return true
}
