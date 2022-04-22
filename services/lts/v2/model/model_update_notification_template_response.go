package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateNotificationTemplateResponse struct {

	// 通知规则名称，必填，只含有汉字、数字、字母、下划线、中划线，不能以下划线等特殊符号开头和结尾，长度为 1 - 100，创建后不可修改
	Name *string `json:"name,omitempty"`

	// 保留字段，非必填，只支持sms（短信），dingding（钉钉），wechat（企业微信），email（邮件）和webhook（网络钩子）
	Type *[]string `json:"type,omitempty"`

	// 模板描述，必填，只含有汉字、数字、字母、下划线不能以下划线开头和结尾，长度为0--1024
	Desc *string `json:"desc,omitempty"`

	// 模板来源，目前必填为LTS，否则会筛选不出来
	Source *string `json:"source,omitempty"`

	// 语言，必填，目前可填zh-cn和en-us
	Locale *UpdateNotificationTemplateResponseLocale `json:"locale,omitempty"`

	// 模板正文，为一个数组
	Templates      *[]SubTemplate `json:"templates,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateNotificationTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateNotificationTemplateResponse", string(data)}, " ")
}

type UpdateNotificationTemplateResponseLocale struct {
	value string
}

type UpdateNotificationTemplateResponseLocaleEnum struct {
	ZH_CN UpdateNotificationTemplateResponseLocale
	EN_US UpdateNotificationTemplateResponseLocale
}

func GetUpdateNotificationTemplateResponseLocaleEnum() UpdateNotificationTemplateResponseLocaleEnum {
	return UpdateNotificationTemplateResponseLocaleEnum{
		ZH_CN: UpdateNotificationTemplateResponseLocale{
			value: "zh-cn",
		},
		EN_US: UpdateNotificationTemplateResponseLocale{
			value: "en-us",
		},
	}
}

func (c UpdateNotificationTemplateResponseLocale) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNotificationTemplateResponseLocale) UnmarshalJSON(b []byte) error {
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
