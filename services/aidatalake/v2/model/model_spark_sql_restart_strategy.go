package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SparkSqlRestartStrategy **参数解释**：集群重启策略，用于指定重启方式。 **约束限制**：不涉及。 **取值范围**：         - FORCE：强制重启，不等待正在运行的作业完成。 - GRACEFUL：优雅重启，等待所有正在运行的作业完成后再重启。 **默认取值**：不涉及。
type SparkSqlRestartStrategy struct {
	value string
}

type SparkSqlRestartStrategyEnum struct {
	FORCE    SparkSqlRestartStrategy
	GRACEFUL SparkSqlRestartStrategy
}

func GetSparkSqlRestartStrategyEnum() SparkSqlRestartStrategyEnum {
	return SparkSqlRestartStrategyEnum{
		FORCE: SparkSqlRestartStrategy{
			value: "FORCE",
		},
		GRACEFUL: SparkSqlRestartStrategy{
			value: "GRACEFUL",
		},
	}
}

func (c SparkSqlRestartStrategy) Value() string {
	return c.value
}

func (c SparkSqlRestartStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SparkSqlRestartStrategy) UnmarshalJSON(b []byte) error {
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
