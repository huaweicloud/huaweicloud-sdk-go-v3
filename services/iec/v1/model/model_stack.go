package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘资源组对象
type Stack struct {

	// 边缘资源组名称。 取值范围：只能由中文字符、大小写英文字母、数字及中划线、下划线组成，且长度为[1-48]个字符。
	Name string `json:"name"`

	// 边缘业务的堆栈，即为资源组。
	Resources []Resource `json:"resources"`
}

func (o Stack) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Stack struct{}"
	}

	return strings.Join([]string{"Stack", string(data)}, " ")
}
