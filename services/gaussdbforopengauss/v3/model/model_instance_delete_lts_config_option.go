package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type InstanceDeleteLtsConfigOption struct {

	// **参数解释**: 实例ID。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 日志类型。 **约束限制**: 不涉及。 **取值范围**: - audit_log：审计日志。  **默认取值**: 不涉及。
	LogType InstanceDeleteLtsConfigOptionLogType `json:"log_type"`
}

func (o InstanceDeleteLtsConfigOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDeleteLtsConfigOption struct{}"
	}

	return strings.Join([]string{"InstanceDeleteLtsConfigOption", string(data)}, " ")
}

type InstanceDeleteLtsConfigOptionLogType struct {
	value string
}

type InstanceDeleteLtsConfigOptionLogTypeEnum struct {
	AUDIT_LOG InstanceDeleteLtsConfigOptionLogType
}

func GetInstanceDeleteLtsConfigOptionLogTypeEnum() InstanceDeleteLtsConfigOptionLogTypeEnum {
	return InstanceDeleteLtsConfigOptionLogTypeEnum{
		AUDIT_LOG: InstanceDeleteLtsConfigOptionLogType{
			value: "audit_log",
		},
	}
}

func (c InstanceDeleteLtsConfigOptionLogType) Value() string {
	return c.value
}

func (c InstanceDeleteLtsConfigOptionLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceDeleteLtsConfigOptionLogType) UnmarshalJSON(b []byte) error {
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
