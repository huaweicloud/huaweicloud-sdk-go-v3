package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SwitchWdrSnapshotStatusRequest Request Object
type SwitchWdrSnapshotStatusRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *SwitchWdrSnapshotStatusRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *SwitchWdrSnapshotRequestBody `json:"body,omitempty"`
}

func (o SwitchWdrSnapshotStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchWdrSnapshotStatusRequest struct{}"
	}

	return strings.Join([]string{"SwitchWdrSnapshotStatusRequest", string(data)}, " ")
}

type SwitchWdrSnapshotStatusRequestXLanguage struct {
	value string
}

type SwitchWdrSnapshotStatusRequestXLanguageEnum struct {
	ZH_CN SwitchWdrSnapshotStatusRequestXLanguage
	EN_US SwitchWdrSnapshotStatusRequestXLanguage
}

func GetSwitchWdrSnapshotStatusRequestXLanguageEnum() SwitchWdrSnapshotStatusRequestXLanguageEnum {
	return SwitchWdrSnapshotStatusRequestXLanguageEnum{
		ZH_CN: SwitchWdrSnapshotStatusRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: SwitchWdrSnapshotStatusRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c SwitchWdrSnapshotStatusRequestXLanguage) Value() string {
	return c.value
}

func (c SwitchWdrSnapshotStatusRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchWdrSnapshotStatusRequestXLanguage) UnmarshalJSON(b []byte) error {
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
