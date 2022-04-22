package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 执行部署任务时传递的动态参数
type KeyValueDo struct {

	// 执行部署任务时传递的参数名称
	Name *string `json:"name,omitempty"`

	// 执行部署任务时传递的参数值
	Value *string `json:"value,omitempty"`

	// 参数值为枚举类型时，返回可选值列表
	Limits *[]ParamTypeLimits `json:"limits,omitempty"`
}

func (o KeyValueDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyValueDo struct{}"
	}

	return strings.Join([]string{"KeyValueDo", string(data)}, " ")
}
