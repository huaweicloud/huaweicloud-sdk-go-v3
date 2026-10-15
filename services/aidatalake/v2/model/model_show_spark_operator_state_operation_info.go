package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSparkOperatorStateOperationInfo **参数解释**：operator操作信息。
type ShowSparkOperatorStateOperationInfo struct {
	Type ShowSparkOperatorStateOperationInfoType `json:"type"`

	// **参数解释**：SparkSql端点名称。 **取值范围**：只能由小写字母、数字及中划线组成，必须以小写字母开头，以小写字母或数字结尾，且长度为1~63个字符。
	EndpointName *string `json:"endpoint_name,omitempty"`

	// **参数解释**：集群重启策略，用于指定重启方式。 **取值范围**：         - FORCE：强制重启，不等待正在运行的作业完成。 - GRACEFUL：优雅重启，等待所有正在运行的作业完成后再重启。
	RestartStrategy *ShowSparkOperatorStateOperationInfoRestartStrategy `json:"restart_strategy,omitempty"`
}

func (o ShowSparkOperatorStateOperationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSparkOperatorStateOperationInfo struct{}"
	}

	return strings.Join([]string{"ShowSparkOperatorStateOperationInfo", string(data)}, " ")
}

type ShowSparkOperatorStateOperationInfoType struct {
	value string
}

type ShowSparkOperatorStateOperationInfoTypeEnum struct {
	RESTART_CLUSTER ShowSparkOperatorStateOperationInfoType
}

func GetShowSparkOperatorStateOperationInfoTypeEnum() ShowSparkOperatorStateOperationInfoTypeEnum {
	return ShowSparkOperatorStateOperationInfoTypeEnum{
		RESTART_CLUSTER: ShowSparkOperatorStateOperationInfoType{
			value: "RESTART_CLUSTER",
		},
	}
}

func (c ShowSparkOperatorStateOperationInfoType) Value() string {
	return c.value
}

func (c ShowSparkOperatorStateOperationInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSparkOperatorStateOperationInfoType) UnmarshalJSON(b []byte) error {
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

type ShowSparkOperatorStateOperationInfoRestartStrategy struct {
	value string
}

type ShowSparkOperatorStateOperationInfoRestartStrategyEnum struct {
	GRACEFUL ShowSparkOperatorStateOperationInfoRestartStrategy
	FORCE    ShowSparkOperatorStateOperationInfoRestartStrategy
}

func GetShowSparkOperatorStateOperationInfoRestartStrategyEnum() ShowSparkOperatorStateOperationInfoRestartStrategyEnum {
	return ShowSparkOperatorStateOperationInfoRestartStrategyEnum{
		GRACEFUL: ShowSparkOperatorStateOperationInfoRestartStrategy{
			value: "GRACEFUL",
		},
		FORCE: ShowSparkOperatorStateOperationInfoRestartStrategy{
			value: "FORCE",
		},
	}
}

func (c ShowSparkOperatorStateOperationInfoRestartStrategy) Value() string {
	return c.value
}

func (c ShowSparkOperatorStateOperationInfoRestartStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSparkOperatorStateOperationInfoRestartStrategy) UnmarshalJSON(b []byte) error {
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
