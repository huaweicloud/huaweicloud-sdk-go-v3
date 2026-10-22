package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SparkSchedulingConfig struct {

	// **参数解释**：端点最大并发作业数量。 **约束限制**：不涉及。 **取值范围**：1~1000。 **默认取值**：200。
	MaxConcurrentJobs *int32 `json:"max_concurrent_jobs,omitempty"`

	// **参数解释**：作业调度策略，支持随机和先进先出。 **约束限制**：不涉及。 **取值范围**： - RANDOM：随机。 - FIFO：先进先出。 **默认取值**：RANDOM。
	SchedulingPolicy *SparkSchedulingConfigSchedulingPolicy `json:"scheduling_policy,omitempty"`

	// **参数解释**：FIFO允许并发时间（秒）。 **约束限制**：不涉及。 **取值范围**：0~3600。 **默认取值**：300。
	FifoToleranceSeconds *int32 `json:"fifo_tolerance_seconds,omitempty"`
}

func (o SparkSchedulingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkSchedulingConfig struct{}"
	}

	return strings.Join([]string{"SparkSchedulingConfig", string(data)}, " ")
}

type SparkSchedulingConfigSchedulingPolicy struct {
	value string
}

type SparkSchedulingConfigSchedulingPolicyEnum struct {
	RANDOM SparkSchedulingConfigSchedulingPolicy
	FIFO   SparkSchedulingConfigSchedulingPolicy
}

func GetSparkSchedulingConfigSchedulingPolicyEnum() SparkSchedulingConfigSchedulingPolicyEnum {
	return SparkSchedulingConfigSchedulingPolicyEnum{
		RANDOM: SparkSchedulingConfigSchedulingPolicy{
			value: "RANDOM",
		},
		FIFO: SparkSchedulingConfigSchedulingPolicy{
			value: "FIFO",
		},
	}
}

func (c SparkSchedulingConfigSchedulingPolicy) Value() string {
	return c.value
}

func (c SparkSchedulingConfigSchedulingPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SparkSchedulingConfigSchedulingPolicy) UnmarshalJSON(b []byte) error {
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
