package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type BatchSetPolicyRequest struct {

	// 请求语言类型。
	XLanguage *BatchSetPolicyRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchSetupSyncPolicyReq `json:"body,omitempty"`
}

func (o BatchSetPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetPolicyRequest struct{}"
	}

	return strings.Join([]string{"BatchSetPolicyRequest", string(data)}, " ")
}

type BatchSetPolicyRequestXLanguage struct {
	value string
}

type BatchSetPolicyRequestXLanguageEnum struct {
	EN_US BatchSetPolicyRequestXLanguage
	ZH_CN BatchSetPolicyRequestXLanguage
}

func GetBatchSetPolicyRequestXLanguageEnum() BatchSetPolicyRequestXLanguageEnum {
	return BatchSetPolicyRequestXLanguageEnum{
		EN_US: BatchSetPolicyRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchSetPolicyRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchSetPolicyRequestXLanguage) Value() string {
	return c.value
}

func (c BatchSetPolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchSetPolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
