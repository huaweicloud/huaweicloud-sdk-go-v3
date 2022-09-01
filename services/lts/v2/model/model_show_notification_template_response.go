package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowNotificationTemplateResponse struct {

	// 通知规则名称，必填，只含有汉字、数字、字母、下划线、中划线，不能以下划线等特殊符号开头和结尾，长度为 1 - 100，创建后不可修改
	Name string `json:"name" xml:"name"`

	// 保留字段，非必填，只支持sms（短信），dingding（钉钉），wechat（企业微信），email（邮件）和webhook（网络钩子）
	Type *[]string `json:"type,omitempty" xml:"type"`

	// 模板描述，必填，只含有汉字、数字、字母、下划线不能以下划线开头和结尾，长度为0--1024
	Desc string `json:"desc" xml:"desc"`

	// 模板来源，目前必填为LTS，否则会筛选不出来
	Source string `json:"source" xml:"source"`

	// 语言，必填，目前可填zh-cn和en-us
	Locale ShowNotificationTemplateResponseLocale `json:"locale" xml:"locale"`

	// 模板正文，为一个数组
	Templates []SubTemplate `json:"templates" xml:"templates"`

	// 创建时间，为毫秒时间戳
	CreateTime int64 `json:"create_time" xml:"create_time"`

	// 更新时间，为毫秒时间戳
	ModifyTime int64 `json:"modify_time" xml:"modify_time"`

	// 项目ID，获取方式请参见：获取账号ID、项目ID、日志组ID、日志流ID（https://support.huaweicloud.com/api-lts/lts_api_0006.html）。
	ProjectId      string `json:"project_id" xml:"project_id"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowNotificationTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNotificationTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowNotificationTemplateResponse", string(data)}, " ")
}

type ShowNotificationTemplateResponseLocale struct {
	value string
}

type ShowNotificationTemplateResponseLocaleEnum struct {
	ZH_CN ShowNotificationTemplateResponseLocale
	EN_US ShowNotificationTemplateResponseLocale
}

func GetShowNotificationTemplateResponseLocaleEnum() ShowNotificationTemplateResponseLocaleEnum {
	return ShowNotificationTemplateResponseLocaleEnum{
		ZH_CN: ShowNotificationTemplateResponseLocale{
			value: "zh-cn",
		},
		EN_US: ShowNotificationTemplateResponseLocale{
			value: "en-us",
		},
	}
}

func (c ShowNotificationTemplateResponseLocale) Value() string {
	return c.value
}

func (c ShowNotificationTemplateResponseLocale) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowNotificationTemplateResponseLocale) UnmarshalJSON(b []byte) error {
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
