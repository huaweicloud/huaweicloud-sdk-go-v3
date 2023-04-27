package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmsTemplateReq struct {

	// 应用主键ID
	AppId string `json:"app_id"`

	// 中括号类型。支持枚举值： 1. CN: 中文类型 2. GB: 英文类型
	Brackets *string `json:"brackets,omitempty"`

	// 地域 1. cn：国内 2. intl：
	Region *string `json:"region,omitempty"`

	// 发送国家id列表，只有地域为国际时，该字段有效
	SendCountry *[]int64 `json:"send_country,omitempty"`

	// 签名主键ID，只有地域为国内时，该字段有效
	SignId *string `json:"sign_id,omitempty"`

	// 模板内容
	TemplateContent string `json:"template_content"`

	// 模板描述
	TemplateDesc *string `json:"template_desc,omitempty"`

	// 模板名称
	TemplateName string `json:"template_name"`

	// 模板类型。只有地域为国内时，该字段有效。支持枚举值： 1. VERIFY_CODE_TYPE: 验证码类 2. PROMOTION_TYPE: 推广类 3. NOTIFY_TYPE: 通知类
	TemplateType string `json:"template_type"`

	// 是否为通用模板 1. 0: 非通用模板 2. 1: 通用模板
	UniversalTemplate *int32 `json:"universal_template,omitempty"`

	// 模板参数
	VariableAttributes *[]SmsTemplateVariableAttrReq `json:"variable_attributes,omitempty"`
}

func (o SmsTemplateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmsTemplateReq struct{}"
	}

	return strings.Join([]string{"SmsTemplateReq", string(data)}, " ")
}
