package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EngineType **参数解释**：引擎类型。 **约束限制**：不涉及。 **取值范围**：   - AuraJob：AuraJob引擎。   - SparkSql：SparkSql引擎。   - SparkJob：SparkJob引擎。   - FlinkJob：FlinkJob引擎。   - RayCluster：RayCluster引擎。   - RayJob：RayJob引擎。   - AuraJobV2：AuraJobV2引擎。   - RayService：RayService引擎。 **默认取值**：不涉及。
type EngineType struct {
	value string
}

type EngineTypeEnum struct {
	AURA_JOB    EngineType
	SPARK_SQL   EngineType
	SPARK_JOB   EngineType
	FLINK_JOB   EngineType
	RAY_CLUSTER EngineType
	RAY_JOB     EngineType
	AURA_JOB_V2 EngineType
	RAY_SERVICE EngineType
}

func GetEngineTypeEnum() EngineTypeEnum {
	return EngineTypeEnum{
		AURA_JOB: EngineType{
			value: "AuraJob",
		},
		SPARK_SQL: EngineType{
			value: "SparkSql",
		},
		SPARK_JOB: EngineType{
			value: "SparkJob",
		},
		FLINK_JOB: EngineType{
			value: "FlinkJob",
		},
		RAY_CLUSTER: EngineType{
			value: "RayCluster",
		},
		RAY_JOB: EngineType{
			value: "RayJob",
		},
		AURA_JOB_V2: EngineType{
			value: "AuraJobV2作业管理",
		},
		RAY_SERVICE: EngineType{
			value: "RayService",
		},
	}
}

func (c EngineType) Value() string {
	return c.value
}

func (c EngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EngineType) UnmarshalJSON(b []byte) error {
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
