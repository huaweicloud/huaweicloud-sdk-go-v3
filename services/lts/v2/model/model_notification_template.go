package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NotificationTemplate struct {

	// 通知规则名称，必填，只含有汉字、数字、字母、下划线、中划线，不能以下划线等特殊符号开头和结尾，长度为 1 - 100，创建后不可修改
	Name string `json:"name"`

	// 保留字段，非必填，只支持sms（短信），dingding（钉钉），wechat（企业微信），email（邮件）和webhook（网络钩子）
	Type *[]string `json:"type,omitempty"`

	// 模板描述，必填，只含有汉字、数字、字母、下划线不能以下划线开头和结尾，长度为0--1024
	Desc *string `json:"desc,omitempty"`

	// 模板来源，目前必填为LTS，否则会筛选不出来
	Source string `json:"source"`

	// 语言，必填，目前可填zh-cn和en-us
	Locale NotificationTemplateLocale `json:"locale"`

	// 模板正文，为一个数组
	Templates []SubTemplate `json:"templates"`

	// 创建时间，为毫秒时间戳
	CreateTime int64 `json:"create_time"`

	// 更新时间，为毫秒时间戳
	ModifyTime int64 `json:"modify_time"`

	// 项目ID，获取方式请参见：获取账号ID、项目ID、日志组ID、日志流ID（https://support.huaweicloud.com/api-lts/lts_api_0006.html）。
	ProjectId string `json:"project_id"`
}

func (o NotificationTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationTemplate struct{}"
	}

	return strings.Join([]string{"NotificationTemplate", string(data)}, " ")
}

type NotificationTemplateLocale struct {
	value string
}

type NotificationTemplateLocaleEnum struct {
	ZH_CN NotificationTemplateLocale
	EN_US NotificationTemplateLocale
}

func GetNotificationTemplateLocaleEnum() NotificationTemplateLocaleEnum {
	return NotificationTemplateLocaleEnum{
		ZH_CN: NotificationTemplateLocale{
			value: "zh-cn",
		},
		EN_US: NotificationTemplateLocale{
			value: "en-us",
		},
	}
}

func (c NotificationTemplateLocale) Value() string {
	return c.value
}

func (c NotificationTemplateLocale) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotificationTemplateLocale) UnmarshalJSON(b []byte) error {
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
