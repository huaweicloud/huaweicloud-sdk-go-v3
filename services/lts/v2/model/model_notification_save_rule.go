package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type NotificationSaveRule struct {
	// 首选项对应的语言

	Language NotificationSaveRuleLanguage `json:"language"`
	// 首选项对应的时区信息

	Timezone *string `json:"timezone,omitempty"`
	// 用户名

	UserName string `json:"user_name"`
	// 主题信息

	Topics []Topics `json:"topics"`
}

func (o NotificationSaveRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationSaveRule struct{}"
	}

	return strings.Join([]string{"NotificationSaveRule", string(data)}, " ")
}

type NotificationSaveRuleLanguage struct {
	value string
}

type NotificationSaveRuleLanguageEnum struct {
	ZH_CN NotificationSaveRuleLanguage
	EN_US NotificationSaveRuleLanguage
}

func GetNotificationSaveRuleLanguageEnum() NotificationSaveRuleLanguageEnum {
	return NotificationSaveRuleLanguageEnum{
		ZH_CN: NotificationSaveRuleLanguage{
			value: "zh-cn",
		},
		EN_US: NotificationSaveRuleLanguage{
			value: "en-us",
		},
	}
}

func (c NotificationSaveRuleLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotificationSaveRuleLanguage) UnmarshalJSON(b []byte) error {
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
