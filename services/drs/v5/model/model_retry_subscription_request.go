package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RetrySubscriptionRequest Request Object
type RetrySubscriptionRequest struct {

	// 请求语言类型。
	XLanguage *RetrySubscriptionRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`
}

func (o RetrySubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetrySubscriptionRequest struct{}"
	}

	return strings.Join([]string{"RetrySubscriptionRequest", string(data)}, " ")
}

type RetrySubscriptionRequestXLanguage struct {
	value string
}

type RetrySubscriptionRequestXLanguageEnum struct {
	EN_US RetrySubscriptionRequestXLanguage
	ZH_CN RetrySubscriptionRequestXLanguage
}

func GetRetrySubscriptionRequestXLanguageEnum() RetrySubscriptionRequestXLanguageEnum {
	return RetrySubscriptionRequestXLanguageEnum{
		EN_US: RetrySubscriptionRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: RetrySubscriptionRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c RetrySubscriptionRequestXLanguage) Value() string {
	return c.value
}

func (c RetrySubscriptionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RetrySubscriptionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
