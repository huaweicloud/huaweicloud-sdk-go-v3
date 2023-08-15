package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateNotificationTemplateResponse Response Object
type CreateNotificationTemplateResponse struct {

	// 通知规则名称，必填，只含有汉字、数字、字母、下划线、中划线，不能以下划线等特殊符号开头和结尾，长度为 1 - 100，创建后不可修改
	Name *string `json:"name,omitempty"`

	// 保留字段，非必填，只支持sms（短信），dingding（钉钉），wechat（企业微信），email（邮件）和webhook（网络钩子）
	Type *[]string `json:"type,omitempty"`

	// 模板描述，必填，只含有汉字、数字、字母、下划线不能以下划线开头和结尾，长度为0--1024
	Desc *string `json:"desc,omitempty"`

	// 模板来源，目前必填为LTS，否则会筛选不出来
	Source *string `json:"source,omitempty"`

	// 语言，必填，目前可填zh-cn和en-us
	Locale *CreateNotificationTemplateResponseLocale `json:"locale,omitempty"`

	// 模板正文，为一个数组
	Templates      *[]SubTemplate `json:"templates,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateNotificationTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateNotificationTemplateResponse", string(data)}, " ")
}

type CreateNotificationTemplateResponseLocale struct {
	value string
}

type CreateNotificationTemplateResponseLocaleEnum struct {
	ZH_CN CreateNotificationTemplateResponseLocale
	EN_US CreateNotificationTemplateResponseLocale
}

func GetCreateNotificationTemplateResponseLocaleEnum() CreateNotificationTemplateResponseLocaleEnum {
	return CreateNotificationTemplateResponseLocaleEnum{
		ZH_CN: CreateNotificationTemplateResponseLocale{
			value: "zh-cn",
		},
		EN_US: CreateNotificationTemplateResponseLocale{
			value: "en-us",
		},
	}
}

func (c CreateNotificationTemplateResponseLocale) Value() string {
	return c.value
}

func (c CreateNotificationTemplateResponseLocale) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNotificationTemplateResponseLocale) UnmarshalJSON(b []byte) error {
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
