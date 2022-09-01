package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PropertyValue struct {

	// 属性名称
	PropertyName *string `json:"property_name,omitempty" xml:"property_name"`

	// 属性值
	Value *interface{} `json:"value,omitempty" xml:"value"`

	// 属性值最后更新时间
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp"`
}

func (o PropertyValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PropertyValue struct{}"
	}

	return strings.Join([]string{"PropertyValue", string(data)}, " ")
}
