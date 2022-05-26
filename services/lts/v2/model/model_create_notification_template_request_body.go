package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建通知规则请求体
type CreateNotificationTemplateRequestBody struct {

	// 通知规则名称，必填，只含有汉字、数字、字母、下划线、中划线，不能以下划线等特殊符号开头和结尾，长度为 1 - 100，创建后不可修改
	Name string `json:"name"`

	// 保留字段，非必填，只支持sms（短信），dingding（钉钉），wechat（企业微信），email（邮件）和webhook（网络钩子）
	Type *[]string `json:"type,omitempty"`

	// 模板描述，必填，只含有汉字、数字、字母、下划线不能以下划线开头和结尾，长度为0--1024
	Desc string `json:"desc"`

	// 模板来源，目前必填为LTS，否则会筛选不出来
	Source string `json:"source"`

	// 语言，必填，目前可填zh-cn和en-us
	Locale CreateNotificationTemplateRequestBodyLocale `json:"locale"`

	// 模板正文，为一个数组
	Templates []SubTemplate `json:"templates"`
}

func (o CreateNotificationTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateNotificationTemplateRequestBody", string(data)}, " ")
}

type CreateNotificationTemplateRequestBodyLocale struct {
	value string
}

type CreateNotificationTemplateRequestBodyLocaleEnum struct {
	ZH_CN CreateNotificationTemplateRequestBodyLocale
	EN_US CreateNotificationTemplateRequestBodyLocale
}

func GetCreateNotificationTemplateRequestBodyLocaleEnum() CreateNotificationTemplateRequestBodyLocaleEnum {
	return CreateNotificationTemplateRequestBodyLocaleEnum{
		ZH_CN: CreateNotificationTemplateRequestBodyLocale{
			value: "zh-cn",
		},
		EN_US: CreateNotificationTemplateRequestBodyLocale{
			value: "en-us",
		},
	}
}

func (c CreateNotificationTemplateRequestBodyLocale) Value() string {
	return c.value
}

func (c CreateNotificationTemplateRequestBodyLocale) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNotificationTemplateRequestBodyLocale) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
