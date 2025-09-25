package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchSetBackupPolicyRequest Request Object
type BatchSetBackupPolicyRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *BatchSetBackupPolicyRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchSetBackupPolicyRequestBody `json:"body,omitempty"`
}

func (o BatchSetBackupPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetBackupPolicyRequest struct{}"
	}

	return strings.Join([]string{"BatchSetBackupPolicyRequest", string(data)}, " ")
}

type BatchSetBackupPolicyRequestXLanguage struct {
	value string
}

type BatchSetBackupPolicyRequestXLanguageEnum struct {
	ZH_CN BatchSetBackupPolicyRequestXLanguage
	EN_US BatchSetBackupPolicyRequestXLanguage
}

func GetBatchSetBackupPolicyRequestXLanguageEnum() BatchSetBackupPolicyRequestXLanguageEnum {
	return BatchSetBackupPolicyRequestXLanguageEnum{
		ZH_CN: BatchSetBackupPolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: BatchSetBackupPolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c BatchSetBackupPolicyRequestXLanguage) Value() string {
	return c.value
}

func (c BatchSetBackupPolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchSetBackupPolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
