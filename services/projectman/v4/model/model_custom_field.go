package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomField 自定义属性
type CustomField struct {

	// 自定义属性名
	Name *string `json:"name,omitempty"`

	// 自定义属性对应的值
	Value *string `json:"value,omitempty"`

	// 自定义属性名
	NewName *string `json:"new_name,omitempty"`
}

func (o CustomField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomField struct{}"
	}

	return strings.Join([]string{"CustomField", string(data)}, " ")
}
