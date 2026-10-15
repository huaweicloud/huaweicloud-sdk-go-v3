package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ResetSubscriptionRequest Request Object
type ResetSubscriptionRequest struct {

	// 请求语言类型。
	XLanguage *ResetSubscriptionRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`
}

func (o ResetSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"ResetSubscriptionRequest", string(data)}, " ")
}

type ResetSubscriptionRequestXLanguage struct {
	value string
}

type ResetSubscriptionRequestXLanguageEnum struct {
	EN_US ResetSubscriptionRequestXLanguage
	ZH_CN ResetSubscriptionRequestXLanguage
}

func GetResetSubscriptionRequestXLanguageEnum() ResetSubscriptionRequestXLanguageEnum {
	return ResetSubscriptionRequestXLanguageEnum{
		EN_US: ResetSubscriptionRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ResetSubscriptionRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ResetSubscriptionRequestXLanguage) Value() string {
	return c.value
}

func (c ResetSubscriptionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResetSubscriptionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
