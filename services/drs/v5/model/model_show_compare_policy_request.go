package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowComparePolicyRequest Request Object
type ShowComparePolicyRequest struct {

	// 请求语言类型。
	XLanguage *ShowComparePolicyRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`
}

func (o ShowComparePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowComparePolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowComparePolicyRequest", string(data)}, " ")
}

type ShowComparePolicyRequestXLanguage struct {
	value string
}

type ShowComparePolicyRequestXLanguageEnum struct {
	EN_US ShowComparePolicyRequestXLanguage
	ZH_CN ShowComparePolicyRequestXLanguage
}

func GetShowComparePolicyRequestXLanguageEnum() ShowComparePolicyRequestXLanguageEnum {
	return ShowComparePolicyRequestXLanguageEnum{
		EN_US: ShowComparePolicyRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ShowComparePolicyRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ShowComparePolicyRequestXLanguage) Value() string {
	return c.value
}

func (c ShowComparePolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowComparePolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
