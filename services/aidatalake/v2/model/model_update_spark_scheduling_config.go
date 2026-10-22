package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateSparkSchedulingConfig struct {

	// **参数解释**：端点最大并发作业数量。 **约束限制**：不涉及。 **取值范围**：1~1000。 **默认取值**：不涉及。
	MaxConcurrentJobs *int32 `json:"max_concurrent_jobs,omitempty"`

	// **参数解释**：作业调度策略，支持随机和先进先出。 **约束限制**：不涉及。 **取值范围**： - RANDOM：随机。 - FIFO：先进先出。 **默认取值**：不涉及。
	SchedulingPolicy *UpdateSparkSchedulingConfigSchedulingPolicy `json:"scheduling_policy,omitempty"`

	// **参数解释**：FIFO允许并发时间（秒）。 **约束限制**：不涉及。 **取值范围**：0~3600。 **默认取值**：不涉及。
	FifoToleranceSeconds *int32 `json:"fifo_tolerance_seconds,omitempty"`
}

func (o UpdateSparkSchedulingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSparkSchedulingConfig struct{}"
	}

	return strings.Join([]string{"UpdateSparkSchedulingConfig", string(data)}, " ")
}

type UpdateSparkSchedulingConfigSchedulingPolicy struct {
	value string
}

type UpdateSparkSchedulingConfigSchedulingPolicyEnum struct {
	RANDOM UpdateSparkSchedulingConfigSchedulingPolicy
	FIFO   UpdateSparkSchedulingConfigSchedulingPolicy
}

func GetUpdateSparkSchedulingConfigSchedulingPolicyEnum() UpdateSparkSchedulingConfigSchedulingPolicyEnum {
	return UpdateSparkSchedulingConfigSchedulingPolicyEnum{
		RANDOM: UpdateSparkSchedulingConfigSchedulingPolicy{
			value: "RANDOM",
		},
		FIFO: UpdateSparkSchedulingConfigSchedulingPolicy{
			value: "FIFO",
		},
	}
}

func (c UpdateSparkSchedulingConfigSchedulingPolicy) Value() string {
	return c.value
}

func (c UpdateSparkSchedulingConfigSchedulingPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSparkSchedulingConfigSchedulingPolicy) UnmarshalJSON(b []byte) error {
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
