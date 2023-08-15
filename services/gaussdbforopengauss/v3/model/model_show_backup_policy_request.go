package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowBackupPolicyRequest Request Object
type ShowBackupPolicyRequest struct {

	// 语言。默认值：en-us。
	XLanguage *ShowBackupPolicyRequestXLanguage `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowBackupPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowBackupPolicyRequest", string(data)}, " ")
}

type ShowBackupPolicyRequestXLanguage struct {
	value string
}

type ShowBackupPolicyRequestXLanguageEnum struct {
	ZH_CN ShowBackupPolicyRequestXLanguage
	EN_US ShowBackupPolicyRequestXLanguage
}

func GetShowBackupPolicyRequestXLanguageEnum() ShowBackupPolicyRequestXLanguageEnum {
	return ShowBackupPolicyRequestXLanguageEnum{
		ZH_CN: ShowBackupPolicyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowBackupPolicyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowBackupPolicyRequestXLanguage) Value() string {
	return c.value
}

func (c ShowBackupPolicyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowBackupPolicyRequestXLanguage) UnmarshalJSON(b []byte) error {
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
