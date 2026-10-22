package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateFlinkJob struct {

	// **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：FlinkJob。 **默认取值**：不涉及。
	Type *CreateFlinkJobType `json:"type,omitempty"`

	// **参数解释**：Flink参数配置，Map的key为Flink配置参数名，value为对应的参数值。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	FlinkConfig map[string]string `json:"flink_config,omitempty"`

	StagingConfig *StagingConfig `json:"staging_config"`

	Image *V2ImageBriefInfoAdmin `json:"image,omitempty"`
}

func (o CreateFlinkJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkJob struct{}"
	}

	return strings.Join([]string{"CreateFlinkJob", string(data)}, " ")
}

type CreateFlinkJobType struct {
	value string
}

type CreateFlinkJobTypeEnum struct {
	FLINK_JOB CreateFlinkJobType
}

func GetCreateFlinkJobTypeEnum() CreateFlinkJobTypeEnum {
	return CreateFlinkJobTypeEnum{
		FLINK_JOB: CreateFlinkJobType{
			value: "FlinkJob",
		},
	}
}

func (c CreateFlinkJobType) Value() string {
	return c.value
}

func (c CreateFlinkJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFlinkJobType) UnmarshalJSON(b []byte) error {
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
