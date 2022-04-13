package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PropertyValue struct {
	// 属性名称

	PropertyName *string `json:"property_name,omitempty"`
	// 属性值

	Value *interface{} `json:"value,omitempty"`
	// 属性值最后更新时间

	Timestamp *string `json:"timestamp,omitempty"`
}

func (o PropertyValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PropertyValue struct{}"
	}

	return strings.Join([]string{"PropertyValue", string(data)}, " ")
}
