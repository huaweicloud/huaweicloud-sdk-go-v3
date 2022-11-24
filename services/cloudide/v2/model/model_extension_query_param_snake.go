package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtensionQueryParamSnake struct {

	// 过滤字段
	Filters []FilterSnake `json:"filters"`

	// 插件flag;通过传递flag参数来进行过滤或其他操作。flag的基础数字是2\\4\\8\\16;传递的参数只能是这四个数字加法组合而成数字 利用它们之间二进制的运算获取的值进行其他操作.比如6=0110=0010+0100也就是2和4的集合flags
	Flags int32 `json:"flags"`
}

func (o ExtensionQueryParamSnake) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionQueryParamSnake struct{}"
	}

	return strings.Join([]string{"ExtensionQueryParamSnake", string(data)}, " ")
}
