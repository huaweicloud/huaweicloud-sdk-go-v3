package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AgentStatusInfo struct {

	// **参数解释**: 机器id **取值范围**: 1到64个字符的字符串，且只包含字母、数字和连字符
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: uniagent运行状态 **取值范围**: - none: 未安装 - running: 运行中 - silent: 静默状态，用于大规模插件异常时，紧急规避的一种措施，现象是kill掉telescope，只保留uniagent的心跳功能 - unknown: 心跳故障，不上报心跳数据，属于连接丢失故障
	UniagentStatus *AgentStatusInfoUniagentStatus `json:"uniagent_status,omitempty"`

	// **参数解释**: 插件信息列表 **取值范围**: 数组长度为[1,10]
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
