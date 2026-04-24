package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSubscriptionDetailRequest Request Object
type ShowSubscriptionDetailRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *ShowSubscriptionDetailRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowSubscriptionDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubscriptionDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowSubscriptionDetailRequest", string(data)}, " ")
}

type ShowSubscriptionDetailRequestXLanguage struct {
	value string
}

type ShowSubscriptionDetailRequestXLanguageEnum struct {
	EN_US ShowSubscriptionDetailRequestXLanguage
	ZH_CN ShowSubscriptionDetailRequestXLanguage
}

func GetShowSubscriptionDetailRequestXLanguageEnum() ShowSubscriptionDetailRequestXLanguageEnum {
	return ShowSubscriptionDetailRequestXLanguageEnum{
		EN_US: ShowSubscriptionDetailRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowSubscriptionDetailRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowSubscriptionDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSubscriptionDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSubscriptionDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
