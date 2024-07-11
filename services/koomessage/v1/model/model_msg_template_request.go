package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MsgTemplateRequest 创建短信模板请求体。
type MsgTemplateRequest struct {

	// 应用ID，默认取签名所属的应用ID。
	AppId *string `json:"app_id,omitempty"`

	// 签名ID。
	SignatureId string `json:"signature_id"`

	// 模板内容。
	TemplateContent string `json:"template_content"`

	// 模板描述。
	TemplateDesc *string `json:"template_desc,omitempty"`

	// 模板名称。
	TemplateName string `json:"template_name"`

	// 模板类型。默认取所属签名的签名类型。PROMOTION_TYPE：营销类，NOTIFY_TYPE：通知类。
	TemplateType *MsgTemplateRequestTemplateType `json:"template_type,omitempty"`

	// 是否为通用模板(暂不支持通用模板)。0：非通用模板，1：通用模板。
	UniversalTemplate *MsgTemplateRequestUniversalTemplate `json:"universal_template,omitempty"`

	// 模板参数。
	VariableAttributes []VariableAttributes `json:"variable_attributes"`
}

func (o MsgTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MsgTemplateRequest struct{}"
	}

	return strings.Join([]string{"MsgTemplateRequest", string(data)}, " ")
}

type MsgTemplateRequestTemplateType struct {
	value string
}

type MsgTemplateRequestTemplateTypeEnum struct {
	PROMOTION_TYPE MsgTemplateRequestTemplateType
	NOTIFY_TYPE    MsgTemplateRequestTemplateType
}

func GetMsgTemplateRequestTemplateTypeEnum() MsgTemplateRequestTemplateTypeEnum {
	return MsgTemplateRequestTemplateTypeEnum{
		PROMOTION_TYPE: MsgTemplateRequestTemplateType{
			value: "PROMOTION_TYPE",
		},
		NOTIFY_TYPE: MsgTemplateRequestTemplateType{
			value: "NOTIFY_TYPE",
		},
	}
}

func (c MsgTemplateRequestTemplateType) Value() string {
	return c.value
}

func (c MsgTemplateRequestTemplateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MsgTemplateRequestTemplateType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type MsgTemplateRequestUniversalTemplate struct {
	value string
}

type MsgTemplateRequestUniversalTemplateEnum struct {
	E_0 MsgTemplateRequestUniversalTemplate
	E_1 MsgTemplateRequestUniversalTemplate
}

func GetMsgTemplateRequestUniversalTemplateEnum() MsgTemplateRequestUniversalTemplateEnum {
	return MsgTemplateRequestUniversalTemplateEnum{
		E_0: MsgTemplateRequestUniversalTemplate{
			value: "0",
		},
		E_1: MsgTemplateRequestUniversalTemplate{
			value: "1",
		},
	}
}

func (c MsgTemplateRequestUniversalTemplate) Value() string {
	return c.value
}

func (c MsgTemplateRequestUniversalTemplate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MsgTemplateRequestUniversalTemplate) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
