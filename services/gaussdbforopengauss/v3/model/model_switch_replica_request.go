package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SwitchReplicaRequest Request Object
type SwitchReplicaRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**:   - zh-cn   - en-us  **默认取值**: en-us
	XLanguage *SwitchReplicaRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *SwitchReplicaRequestBody `json:"body,omitempty"`
}

func (o SwitchReplicaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchReplicaRequest struct{}"
	}

	return strings.Join([]string{"SwitchReplicaRequest", string(data)}, " ")
}

type SwitchReplicaRequestXLanguage struct {
	value string
}

type SwitchReplicaRequestXLanguageEnum struct {
	ZH_CN SwitchReplicaRequestXLanguage
	EN_US SwitchReplicaRequestXLanguage
}

func GetSwitchReplicaRequestXLanguageEnum() SwitchReplicaRequestXLanguageEnum {
	return SwitchReplicaRequestXLanguageEnum{
		ZH_CN: SwitchReplicaRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: SwitchReplicaRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c SwitchReplicaRequestXLanguage) Value() string {
	return c.value
}

func (c SwitchReplicaRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchReplicaRequestXLanguage) UnmarshalJSON(b []byte) error {
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
