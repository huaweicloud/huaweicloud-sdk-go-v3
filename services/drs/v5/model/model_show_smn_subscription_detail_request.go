package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSmnSubscriptionDetailRequest Request Object
type ShowSmnSubscriptionDetailRequest struct {

	// 请求语言类型。
	XLanguage *ShowSmnSubscriptionDetailRequestXLanguage `json:"X-Language,omitempty"`

	// 任务id。
	JobId string `json:"job_id"`
}

func (o ShowSmnSubscriptionDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSmnSubscriptionDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowSmnSubscriptionDetailRequest", string(data)}, " ")
}

type ShowSmnSubscriptionDetailRequestXLanguage struct {
	value string
}

type ShowSmnSubscriptionDetailRequestXLanguageEnum struct {
	EN_US ShowSmnSubscriptionDetailRequestXLanguage
	ZH_CN ShowSmnSubscriptionDetailRequestXLanguage
}

func GetShowSmnSubscriptionDetailRequestXLanguageEnum() ShowSmnSubscriptionDetailRequestXLanguageEnum {
	return ShowSmnSubscriptionDetailRequestXLanguageEnum{
		EN_US: ShowSmnSubscriptionDetailRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowSmnSubscriptionDetailRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowSmnSubscriptionDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ShowSmnSubscriptionDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSmnSubscriptionDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
