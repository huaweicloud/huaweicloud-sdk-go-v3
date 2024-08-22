package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NoticeRuleNotification struct {

	// 通知的协议类型，包括短信，邮件，企业微信群消息等。
	Protocol NoticeRuleNotificationProtocol `json:"protocol"`

	// 通知的终端地址。 email协议，接入点必须是邮件地址。 sms协议，接入点必须是一个电话号码。 wechat协议，参考https://support.huaweicloud.com/smn_faq/smn_faq_0027.html获取订阅终端， 企业微信群消息为SMN服务公测功能，需提交工单申请开通。
	Endpoint string `json:"endpoint"`

	// 通知的模板语言。 ZH，中文。 EN，英文。
	Template *NoticeRuleNotificationTemplate `json:"template,omitempty"`
}

func (o NoticeRuleNotification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoticeRuleNotification struct{}"
	}

	return strings.Join([]string{"NoticeRuleNotification", string(data)}, " ")
}

type NoticeRuleNotificationProtocol struct {
	value string
}

type NoticeRuleNotificationProtocolEnum struct {
	EMAIL  NoticeRuleNotificationProtocol
	SMS    NoticeRuleNotificationProtocol
	WECHAT NoticeRuleNotificationProtocol
}

func GetNoticeRuleNotificationProtocolEnum() NoticeRuleNotificationProtocolEnum {
	return NoticeRuleNotificationProtocolEnum{
		EMAIL: NoticeRuleNotificationProtocol{
			value: "email",
		},
		SMS: NoticeRuleNotificationProtocol{
			value: "sms",
		},
		WECHAT: NoticeRuleNotificationProtocol{
			value: "wechat",
		},
	}
}

func (c NoticeRuleNotificationProtocol) Value() string {
	return c.value
}

func (c NoticeRuleNotificationProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NoticeRuleNotificationProtocol) UnmarshalJSON(b []byte) error {
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

type NoticeRuleNotificationTemplate struct {
	value string
}

type NoticeRuleNotificationTemplateEnum struct {
	ZH NoticeRuleNotificationTemplate
	EN NoticeRuleNotificationTemplate
}

func GetNoticeRuleNotificationTemplateEnum() NoticeRuleNotificationTemplateEnum {
	return NoticeRuleNotificationTemplateEnum{
		ZH: NoticeRuleNotificationTemplate{
			value: "ZH",
		},
		EN: NoticeRuleNotificationTemplate{
			value: "EN",
		},
	}
}

func (c NoticeRuleNotificationTemplate) Value() string {
	return c.value
}

func (c NoticeRuleNotificationTemplate) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NoticeRuleNotificationTemplate) UnmarshalJSON(b []byte) error {
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
