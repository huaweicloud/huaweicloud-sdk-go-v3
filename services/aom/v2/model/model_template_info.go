package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateInfo struct {

	// 变量名称。
	Name *string `json:"name,omitempty"`

	// 变量类型。
	Type *string `json:"type,omitempty"`

	// 变量值。
	Query *string `json:"query,omitempty"`

	// 变量描述。
	Description *string `json:"description,omitempty"`
}

func (o TemplateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateInfo struct{}"
	}

	return strings.Join([]string{"TemplateInfo", string(data)}, " ")
}
