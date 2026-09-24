package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Attribute struct {

	// 属性编码
	Code *string `json:"code,omitempty"`

	// 属性名称
	Name *string `json:"name,omitempty"`

	// 属性取值
	Value *string `json:"value,omitempty"`

	// 属性单位
	Unit *string `json:"unit,omitempty"`

	LinearRange *LinearRange `json:"linear_range,omitempty"`
}

func (o Attribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Attribute struct{}"
	}

	return strings.Join([]string{"Attribute", string(data)}, " ")
}
