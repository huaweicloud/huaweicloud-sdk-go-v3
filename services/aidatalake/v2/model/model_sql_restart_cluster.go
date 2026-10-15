package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SqlRestartCluster **参数解释**：重启集群操作的详情。
type SqlRestartCluster struct {

	// **参数解释**：作业类型。 **约束限制**：固定值为 RESTART_CLUSTER。 **取值范围**：不涉及。 **默认取值**：RESTART_CLUSTER。
	Type *string `json:"type,omitempty"`

	// **参数解释**：SparkSql端点名称。 **取值范围**：只能由小写字母、数字及中划线组成，必须以小写字母开头，以小写字母或数字结尾，且长度为1~63个字符。
	EndpointName *string `json:"endpoint_name,omitempty"`

	// **参数解释**：集群重启策略，用于指定重启方式。 **取值范围**：         - FORCE：强制重启，不等待正在运行的作业完成。 - GRACEFUL：优雅重启，等待所有正在运行的作业完成后再重启。
	RestartStrategy *SqlRestartClusterRestartStrategy `json:"restart_strategy,omitempty"`
}

func (o SqlRestartCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlRestartCluster struct{}"
	}

	return strings.Join([]string{"SqlRestartCluster", string(data)}, " ")
}

type SqlRestartClusterRestartStrategy struct {
	value string
}

type SqlRestartClusterRestartStrategyEnum struct {
	GRACEFUL SqlRestartClusterRestartStrategy
	FORCE    SqlRestartClusterRestartStrategy
}

func GetSqlRestartClusterRestartStrategyEnum() SqlRestartClusterRestartStrategyEnum {
	return SqlRestartClusterRestartStrategyEnum{
		GRACEFUL: SqlRestartClusterRestartStrategy{
			value: "GRACEFUL",
		},
		FORCE: SqlRestartClusterRestartStrategy{
			value: "FORCE",
		},
	}
}

func (c SqlRestartClusterRestartStrategy) Value() string {
	return c.value
}

func (c SqlRestartClusterRestartStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlRestartClusterRestartStrategy) UnmarshalJSON(b []byte) error {
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
