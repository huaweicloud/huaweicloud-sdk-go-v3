package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type InstanceSaveLtsConfigOption struct {

	// **参数解释**: 实例ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 日志类型。 **约束限制**: 不涉及。 **取值范围**: - audit_log：审计日志。  **默认取值**: 不涉及。
	LogType InstanceSaveLtsConfigOptionLogType `json:"log_type"`

	// **参数解释**: LTS日志组ID。 **约束限制**: 不涉及。 **取值范围**: 可通过云日志服务-\"查询账号下所有日志组\"API接口获取。 **默认取值**: 不涉及。
	LtsGroupId string `json:"lts_group_id"`

	// **参数解释**: LTS日志流ID。 **约束限制**: 不涉及。 **取值范围**: 可通过云日志服务-”查询指定日志组下的所有日志流“API接口获取。 **默认取值**: 不涉及。
	LtsStreamId string `json:"lts_stream_id"`
}

func (o InstanceSaveLtsConfigOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceSaveLtsConfigOption struct{}"
	}

	return strings.Join([]string{"InstanceSaveLtsConfigOption", string(data)}, " ")
}

type InstanceSaveLtsConfigOptionLogType struct {
	value string
}

type InstanceSaveLtsConfigOptionLogTypeEnum struct {
	AUDIT_LOG InstanceSaveLtsConfigOptionLogType
}

func GetInstanceSaveLtsConfigOptionLogTypeEnum() InstanceSaveLtsConfigOptionLogTypeEnum {
	return InstanceSaveLtsConfigOptionLogTypeEnum{
		AUDIT_LOG: InstanceSaveLtsConfigOptionLogType{
			value: "audit_log",
		},
	}
}

func (c InstanceSaveLtsConfigOptionLogType) Value() string {
	return c.value
}

func (c InstanceSaveLtsConfigOptionLogType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceSaveLtsConfigOptionLogType) UnmarshalJSON(b []byte) error {
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
