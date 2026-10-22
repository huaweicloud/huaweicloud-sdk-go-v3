package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateSparkJob **参数解释**：SparkJob引擎配置。
type CreateSparkJob struct {

	// **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：SparkJob。 **默认取值**：不涉及。
	Type *CreateSparkJobType `json:"type,omitempty"`

	// **参数解释**：Spark参数配置，Map的key为Spark配置参数名，value为对应的参数值。 **约束限制**：最大数量为100。 **取值范围**：不涉及。 **默认取值**：不涉及。
	SparkConfig map[string]string `json:"spark_config,omitempty"`

	SchedulingConfig *SparkSchedulingConfig `json:"scheduling_config,omitempty"`

	StagingConfig *StagingConfig `json:"staging_config"`

	Image *V2ImageBriefInfoAdmin `json:"image,omitempty"`
}

func (o CreateSparkJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSparkJob struct{}"
	}

	return strings.Join([]string{"CreateSparkJob", string(data)}, " ")
}

type CreateSparkJobType struct {
	value string
}

type CreateSparkJobTypeEnum struct {
	SPARK_JOB CreateSparkJobType
}

func GetCreateSparkJobTypeEnum() CreateSparkJobTypeEnum {
	return CreateSparkJobTypeEnum{
		SPARK_JOB: CreateSparkJobType{
			value: "SparkJob",
		},
	}
}

func (c CreateSparkJobType) Value() string {
	return c.value
}

func (c CreateSparkJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateSparkJobType) UnmarshalJSON(b []byte) error {
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
