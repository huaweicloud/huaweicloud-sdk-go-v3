package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NoticeTypeNotificationTemplateList struct {

	// 通知类型, sms(短信)/email/http(s)/smn/welink/dingding(钉钉)/wechat(微信)/feishu(飞书)/wecomgroup(微信企业版)
	NoticeType NoticeTypeNotificationTemplateListNoticeType `json:"notice_type"`

	// **参数解释**： 通知模板id **约束限制**： 不涉及。 **取值范围**： 长度为[3,64]个字符。 **默认取值**： 不涉及
	AlarmNotificationTemplateId string `json:"alarm_notification_template_id"`

	// **参数解释**： 通知模板名 **约束限制**： 不涉及。 **取值范围**： 长度为[1,200]个字符。 **默认取值**： 不涉及
	AlarmNotificationTemplateName *string `json:"alarm_notification_template_name,omitempty"`
}

func (o NoticeTypeNotificationTemplateList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoticeTypeNotificationTemplateList struct{}"
	}

	return strings.Join([]string{"NoticeTypeNotificationTemplateList", string(data)}, " ")
}

type NoticeTypeNotificationTemplateListNoticeType struct {
	value string
}

type NoticeTypeNotificationTemplateListNoticeTypeEnum struct {
	SMS        NoticeTypeNotificationTemplateListNoticeType
	EMAIL      NoticeTypeNotificationTemplateListNoticeType
	HTTP_S     NoticeTypeNotificationTemplateListNoticeType
	SMN        NoticeTypeNotificationTemplateListNoticeType
	WELINK     NoticeTypeNotificationTemplateListNoticeType
	DINGDING   NoticeTypeNotificationTemplateListNoticeType
	WECHAT     NoticeTypeNotificationTemplateListNoticeType
	FEISHU     NoticeTypeNotificationTemplateListNoticeType
	WECOMGROUP NoticeTypeNotificationTemplateListNoticeType
}

func GetNoticeTypeNotificationTemplateListNoticeTypeEnum() NoticeTypeNotificationTemplateListNoticeTypeEnum {
	return NoticeTypeNotificationTemplateListNoticeTypeEnum{
		SMS: NoticeTypeNotificationTemplateListNoticeType{
			value: "sms",
		},
		EMAIL: NoticeTypeNotificationTemplateListNoticeType{
			value: "email",
		},
		HTTP_S: NoticeTypeNotificationTemplateListNoticeType{
			value: "http(s)",
		},
		SMN: NoticeTypeNotificationTemplateListNoticeType{
			value: "smn",
		},
		WELINK: NoticeTypeNotificationTemplateListNoticeType{
			value: "welink",
		},
		DINGDING: NoticeTypeNotificationTemplateListNoticeType{
			value: "dingding",
		},
		WECHAT: NoticeTypeNotificationTemplateListNoticeType{
			value: "wechat",
		},
		FEISHU: NoticeTypeNotificationTemplateListNoticeType{
			value: "feishu",
		},
		WECOMGROUP: NoticeTypeNotificationTemplateListNoticeType{
			value: "wecomgroup",
		},
	}
}

func (c NoticeTypeNotificationTemplateListNoticeType) Value() string {
	return c.value
}

func (c NoticeTypeNotificationTemplateListNoticeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NoticeTypeNotificationTemplateListNoticeType) UnmarshalJSON(b []byte) error {
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
