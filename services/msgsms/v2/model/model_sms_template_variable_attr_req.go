package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmsTemplateVariableAttrReq struct {

	// 变量说明，当变量类型为TEXT时，必填
	VariableDesc *string `json:"variable_desc,omitempty"`

	// 变量索引，对应模板内容变量索引
	VariableIndex int32 `json:"variable_index"`

	// 变量类型，目前支持：PHONE|CHARDIGIT|DATETIME|MONEY|TEXT|NEWTEXT
	VariableType string `json:"variable_type"`
}

func (o SmsTemplateVariableAttrReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmsTemplateVariableAttrReq struct{}"
	}

	return strings.Join([]string{"SmsTemplateVariableAttrReq", string(data)}, " ")
}
