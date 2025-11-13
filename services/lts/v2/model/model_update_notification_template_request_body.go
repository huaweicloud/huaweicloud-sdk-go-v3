package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateNotificationTemplateRequestBody struct {

	// 通知规则名称，必填，只含有汉字、数字、字母、下划线、中划线，不能以下划线等特殊符号开头和结尾，长度为 1 - 100，创建后不可修改
	Name string `json:"name"`

	// 保留字段，非必填
	Type *[]string `json:"type,omitempty"`

	// 模板描述，非必填，只含有汉字、数字、字母、下划线不能以下划线开头和结尾，长度为0--1024
	Desc *string `json:"desc,omitempty"`

	// 模板来源，目前必填为LTS，否则会筛选不出来
	Source string `json:"source"`

	// 语言，必填，目前可填zh-cn和en-us
	Locale UpdateNotificationTemplateRequestBodyLocale `json:"locale"`

	// 模板正文，为一个数组
	Templates []UpdateSubTemplate `json:"templates"`
}

func (o UpdateNotificationTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNotificationTemplateRequestBody", string(data)}, " ")
}

type UpdateNotificationTemplateRequestBodyLocale struct {
	value string
}

type UpdateNotificationTemplateRequestBodyLocaleEnum struct {
	ZH_CN UpdateNotificationTemplateRequestBodyLocale
	EN_US UpdateNotificationTemplateRequestBodyLocale
}

func GetUpdateNotificationTemplateRequestBodyLocaleEnum() UpdateNotificationTemplateRequestBodyLocaleEnum {
	return UpdateNotificationTemplateRequestBodyLocaleEnum{
		ZH_CN: UpdateNotificationTemplateRequestBodyLocale{
			value: "zh-cn",
		},
		EN_US: UpdateNotificationTemplateRequestBodyLocale{
			value: "en-us",
		},
	}
}

func (c UpdateNotificationTemplateRequestBodyLocale) Value() string {
	return c.value
}

func (c UpdateNotificationTemplateRequestBodyLocale) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNotificationTemplateRequestBodyLocale) UnmarshalJSON(b []byte) error {
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
