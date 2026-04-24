package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteSubscriptionRequest Request Object
type DeleteSubscriptionRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *DeleteSubscriptionRequestXLanguage `json:"X-Language,omitempty"`
}

func (o DeleteSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubscriptionRequest", string(data)}, " ")
}

type DeleteSubscriptionRequestXLanguage struct {
	value string
}

type DeleteSubscriptionRequestXLanguageEnum struct {
	EN_US DeleteSubscriptionRequestXLanguage
	ZH_CN DeleteSubscriptionRequestXLanguage
}

func GetDeleteSubscriptionRequestXLanguageEnum() DeleteSubscriptionRequestXLanguageEnum {
	return DeleteSubscriptionRequestXLanguageEnum{
		EN_US: DeleteSubscriptionRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: DeleteSubscriptionRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c DeleteSubscriptionRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteSubscriptionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteSubscriptionRequestXLanguage) UnmarshalJSON(b []byte) error {
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
