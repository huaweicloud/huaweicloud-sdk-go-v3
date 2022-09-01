package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RawValue struct {

	// 属性名称
	PropertyName *string `json:"property_name,omitempty" xml:"property_name"`

	// 资产属性的历史值序列，示例：[1,2]
	Values *[]interface{} `json:"values,omitempty" xml:"values"`
}

func (o RawValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RawValue struct{}"
	}

	return strings.Join([]string{"RawValue", string(data)}, " ")
}
