package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateComparePolicyRequest Request Object
type UpdateComparePolicyRequest struct {

	// 请求语言类型。
	XLanguage *UpdateComparePolicyRequestXLanguage `json:"X-Language,omitempty"`

	// 任务ID。
	JobId string `json:"job_id"`

	Body *ModifyComparePolicyReq `json:"body,omitempty"`
}

func (o UpdateComparePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComparePolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateComparePolicyRequest", string(data)}, " ")
}

type UpdateComparePolicyRequestXLanguage struct {
	value string
}

type UpdateComparePolicyRequestXLanguageEnum struct {
	EN_US UpdateComparePolicyRequestXLanguage
	ZH_CN UpdateComparePolicyRequestXLanguage
}

func GetUpdateComparePolicyRequestXLanguageEnum() UpdateComparePolicyRequestXLanguageEnum {
	return UpdateComparePolicyRequestXLanguageEnum{
		EN_US: UpdateComparePolicyRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UpdateComparePolicyRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UpdateComparePolicyRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateComparePolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateComparePolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
