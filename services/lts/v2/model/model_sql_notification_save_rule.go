package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SqlNotificationSaveRule struct {

	// 首选项对应的语言
	Language SqlNotificationSaveRuleLanguage `json:"language"`

	// 首选项对应的时区信息
	Timezone *string `json:"timezone,omitempty"`

	// 用户名
	UserName string `json:"user_name"`

	// 主题信息
	Topics []Topics `json:"topics"`

	// 消息模板名称
	TemplateName string `json:"template_name"`
}

func (o SqlNotificationSaveRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlNotificationSaveRule struct{}"
	}

	return strings.Join([]string{"SqlNotificationSaveRule", string(data)}, " ")
}

type SqlNotificationSaveRuleLanguage struct {
	value string
}

type SqlNotificationSaveRuleLanguageEnum struct {
	ZH_CN SqlNotificationSaveRuleLanguage
	EN_US SqlNotificationSaveRuleLanguage
}

func GetSqlNotificationSaveRuleLanguageEnum() SqlNotificationSaveRuleLanguageEnum {
	return SqlNotificationSaveRuleLanguageEnum{
		ZH_CN: SqlNotificationSaveRuleLanguage{
			value: "zh-cn",
		},
		EN_US: SqlNotificationSaveRuleLanguage{
			value: "en-us",
		},
	}
}

func (c SqlNotificationSaveRuleLanguage) Value() string {
	return c.value
}

func (c SqlNotificationSaveRuleLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlNotificationSaveRuleLanguage) UnmarshalJSON(b []byte) error {
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
