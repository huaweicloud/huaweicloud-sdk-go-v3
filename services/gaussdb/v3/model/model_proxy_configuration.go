package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProxyConfiguration struct {

	// 参数名称
	Name *string `json:"name,omitempty"`

	// 参数数据类型
	DataType *string `json:"data_type,omitempty"`

	// 参数父标签类型
	ElemType *string `json:"elem_type,omitempty"`

	// 取值范围
	ValueRange *string `json:"value_range,omitempty"`

	// 参数默认值
	Value *string `json:"value,omitempty"`

	// 参数描述
	Description *string `json:"description,omitempty"`
}

func (o ProxyConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyConfiguration struct{}"
	}

	return strings.Join([]string{"ProxyConfiguration", string(data)}, " ")
}
