package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateFlinkJob struct {

	// **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：FlinkJob。 **默认取值**：不涉及。
	Type *UpdateFlinkJobType `json:"type,omitempty"`

	// **参数解释**：Flink参数配置，Map的key为Flink配置参数名，value为对应的参数值。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	FlinkConfig map[string]string `json:"flink_config,omitempty"`

	Image *V2ImageBriefInfoAdmin `json:"image,omitempty"`
}

func (o UpdateFlinkJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkJob struct{}"
	}

	return strings.Join([]string{"UpdateFlinkJob", string(data)}, " ")
}

type UpdateFlinkJobType struct {
	value string
}

type UpdateFlinkJobTypeEnum struct {
	FLINK_JOB UpdateFlinkJobType
}

func GetUpdateFlinkJobTypeEnum() UpdateFlinkJobTypeEnum {
	return UpdateFlinkJobTypeEnum{
		FLINK_JOB: UpdateFlinkJobType{
			value: "FlinkJob",
		},
	}
}

func (c UpdateFlinkJobType) Value() string {
	return c.value
}

func (c UpdateFlinkJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateFlinkJobType) UnmarshalJSON(b []byte) error {
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
