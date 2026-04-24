package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSubscriptionInfoRequest Request Object
type UpdateSubscriptionInfoRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *UpdateSubscriptionInfoRequestXLanguage `json:"X-Language,omitempty"`

	Body *UpdateSubscriptionInfoReq `json:"body,omitempty"`
}

func (o UpdateSubscriptionInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubscriptionInfoRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubscriptionInfoRequest", string(data)}, " ")
}

type UpdateSubscriptionInfoRequestXLanguage struct {
	value string
}

type UpdateSubscriptionInfoRequestXLanguageEnum struct {
	EN_US UpdateSubscriptionInfoRequestXLanguage
	ZH_CN UpdateSubscriptionInfoRequestXLanguage
}

func GetUpdateSubscriptionInfoRequestXLanguageEnum() UpdateSubscriptionInfoRequestXLanguageEnum {
	return UpdateSubscriptionInfoRequestXLanguageEnum{
		EN_US: UpdateSubscriptionInfoRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UpdateSubscriptionInfoRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UpdateSubscriptionInfoRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateSubscriptionInfoRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSubscriptionInfoRequestXLanguage) UnmarshalJSON(b []byte) error {
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
