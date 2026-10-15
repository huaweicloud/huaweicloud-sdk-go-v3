package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StartSubscriptionRequest Request Object
type StartSubscriptionRequest struct {

	// 请求语言类型。
	XLanguage *StartSubscriptionRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`
}

func (o StartSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"StartSubscriptionRequest", string(data)}, " ")
}

type StartSubscriptionRequestXLanguage struct {
	value string
}

type StartSubscriptionRequestXLanguageEnum struct {
	EN_US StartSubscriptionRequestXLanguage
	ZH_CN StartSubscriptionRequestXLanguage
}

func GetStartSubscriptionRequestXLanguageEnum() StartSubscriptionRequestXLanguageEnum {
	return StartSubscriptionRequestXLanguageEnum{
		EN_US: StartSubscriptionRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: StartSubscriptionRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c StartSubscriptionRequestXLanguage) Value() string {
	return c.value
}

func (c StartSubscriptionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartSubscriptionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
