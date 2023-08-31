package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AgentStatusInfo struct {

	// 机器id
	InstanceId *string `json:"instance_id,omitempty"`

	// uniagent运行状态,none无，running运行中，silent静默中，unknown故障
	UniagentStatus *AgentStatusInfoUniagentStatus `json:"uniagent_status,omitempty"`

	// 插件信息列表
	Extensions *[]ExtensionInfo `json:"extensions,omitempty"`
}

func (o AgentStatusInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentStatusInfo struct{}"
	}

	return strings.Join([]string{"AgentStatusInfo", string(data)}, " ")
}

type AgentStatusInfoUniagentStatus struct {
	value string
}

type AgentStatusInfoUniagentStatusEnum struct {
	NONE    AgentStatusInfoUniagentStatus
	RUNNING AgentStatusInfoUniagentStatus
	SILENT  AgentStatusInfoUniagentStatus
	UNKNOWN AgentStatusInfoUniagentStatus
}

func GetAgentStatusInfoUniagentStatusEnum() AgentStatusInfoUniagentStatusEnum {
	return AgentStatusInfoUniagentStatusEnum{
		NONE: AgentStatusInfoUniagentStatus{
			value: "none",
		},
		RUNNING: AgentStatusInfoUniagentStatus{
			value: "running",
		},
		SILENT: AgentStatusInfoUniagentStatus{
			value: "silent",
		},
		UNKNOWN: AgentStatusInfoUniagentStatus{
			value: "unknown",
		},
	}
}

func (c AgentStatusInfoUniagentStatus) Value() string {
	return c.value
}

func (c AgentStatusInfoUniagentStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgentStatusInfoUniagentStatus) UnmarshalJSON(b []byte) error {
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
