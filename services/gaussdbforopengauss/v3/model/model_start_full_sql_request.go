package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StartFullSqlRequest Request Object
type StartFullSqlRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *StartFullSqlRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为32个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *FullSqlStartRequestBody `json:"body,omitempty"`
}

func (o StartFullSqlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartFullSqlRequest struct{}"
	}

	return strings.Join([]string{"StartFullSqlRequest", string(data)}, " ")
}

type StartFullSqlRequestXLanguage struct {
	value string
}

type StartFullSqlRequestXLanguageEnum struct {
	ZH_CN StartFullSqlRequestXLanguage
	EN_US StartFullSqlRequestXLanguage
}

func GetStartFullSqlRequestXLanguageEnum() StartFullSqlRequestXLanguageEnum {
	return StartFullSqlRequestXLanguageEnum{
		ZH_CN: StartFullSqlRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartFullSqlRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartFullSqlRequestXLanguage) Value() string {
	return c.value
}

func (c StartFullSqlRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartFullSqlRequestXLanguage) UnmarshalJSON(b []byte) error {
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
