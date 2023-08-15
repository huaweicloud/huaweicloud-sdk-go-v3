package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SetBackupPolicyRequest Request Object
type SetBackupPolicyRequest struct {

	// 语言。默认值：en-us。
	XLanguage *SetBackupPolicyRequestXLanguage `json:"X-Language,omitempty"`

	// 实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	Body *SetBackupPolicyRequestBody `json:"body,omitempty"`
}

func (o SetBackupPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetBackupPolicyRequest struct{}"
	}

	return strings.Join([]string{"SetBackupPolicyRequest", string(data)}, " ")
}

type SetBackupPolicyRequestXLanguage struct {
	value string
}

type SetBackupPolicyRequestXLanguageEnum struct {
	ZH_CN SetBackupPolicyRequestXLanguage
	EN_US SetBackupPolicyRequestXLanguage
}

func GetSetBackupPolicyRequestXLanguageEnum() SetBackupPolicyRequestXLanguageEnum {
	return SetBackupPolicyRequestXLanguageEnum{
		ZH_CN: SetBackupPolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: SetBackupPolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c SetBackupPolicyRequestXLanguage) Value() string {
	return c.value
}

func (c SetBackupPolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetBackupPolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
