package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 属性过滤器
type PropertyFilter struct {
	// 过滤属性名称，正则：\"^[a-zA-Z0-9_]{1,64}$\"

	PropertyName string `json:"property_name"`
	// 过滤操作方式,当前仅支持“=”

	Operator string `json:"operator"`
	// 过滤属性值

	Value *interface{} `json:"value"`
}

func (o PropertyFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PropertyFilter struct{}"
	}

	return strings.Join([]string{"PropertyFilter", string(data)}, " ")
}
