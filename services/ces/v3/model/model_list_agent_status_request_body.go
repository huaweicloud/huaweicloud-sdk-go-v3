package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListAgentStatusRequestBody struct {

	// **参数解释**: 机器实例id列表 **约束限制**: 包含的机器实例id最多为2000个，最少为1个
	InstanceIds []string `json:"instance_ids"`

	// **参数解释**: uniagent运行状态，不传值则查出所有状态 **约束限制**: 不涉及。 **取值范围**: - none: 未安装 - running: 运行中 - silent: 静默状态，用于大规模插件异常时，紧急规避的一种措施，现象是kill掉telescope，只保留uniagent的心跳功能 - unknown: 心跳故障，不上报心跳数据，属于连接丢失故障 **默认取值**: 不涉及
	UniagentStatus *ListAgentStatusRequestBodyUniagentStatus `json:"uniagent_status,omitempty"`

	// **参数解释**: 插件名称，不传查所有插件 **约束限制**: 当前仅支持查询telescope插件 **取值范围**: - telescope: 主机监控插件telescope **默认取值**: telescope
	ExtensionName *ListAgentStatusRequestBodyExtensionName `json:"extension_name,omitempty"`

	// **参数解释**: 插件状态，不传查所有状态  **约束限制**: 不涉及 **取值范围**: - none: 未安装 - running: 运行中 - stopped：已停止 - fault: 进程故障，应该运行的插件，没运行，属于客户端故障 - unknown: 心跳故障，不上报心跳数据，属于连接丢失故障 **默认取值**: 不涉及
	ExtensionStatus *ListAgentStatusRequestBodyExtensionStatus `json:"extension_status,omitempty"`
}

func (o ListAgentStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentStatusRequestBody struct{}"
	}

	return strings.Join([]string{"ListAgentStatusRequestBody", string(data)}, " ")
}

type ListAgentStatusRequestBodyUniagentStatus struct {
	value string
}

type ListAgentStatusRequestBodyUniagentStatusEnum struct {
	NONE    ListAgentStatusRequestBodyUniagentStatus
	RUNNING ListAgentStatusRequestBodyUniagentStatus
	SILENT  ListAgentStatusRequestBodyUniagentStatus
	UNKNOWN ListAgentStatusRequestBodyUniagentStatus
}

func GetListAgentStatusRequestBodyUniagentStatusEnum() ListAgentStatusRequestBodyUniagentStatusEnum {
	return ListAgentStatusRequestBodyUniagentStatusEnum{
		NONE: ListAgentStatusRequestBodyUniagentStatus{
			value: "none",
		},
		RUNNING: ListAgentStatusRequestBodyUniagentStatus{
			value: "running",
		},
		SILENT: ListAgentStatusRequestBodyUniagentStatus{
			value: "silent",
		},
		UNKNOWN: ListAgentStatusRequestBodyUniagentStatus{
			value: "unknown",
		},
	}
}

func (c ListAgentStatusRequestBodyUniagentStatus) Value() string {
	return c.value
}

func (c ListAgentStatusRequestBodyUniagentStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgentStatusRequestBodyUniagentStatus) UnmarshalJSON(b []byte) error {
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

type ListAgentStatusRequestBodyExtensionName struct {
	value string
}

type ListAgentStatusRequestBodyExtensionNameEnum struct {
	TELESCOPE ListAgentStatusRequestBodyExtensionName
}

func GetListAgentStatusRequestBodyExtensionNameEnum() ListAgentStatusRequestBodyExtensionNameEnum {
	return ListAgentStatusRequestBodyExtensionNameEnum{
		TELESCOPE: ListAgentStatusRequestBodyExtensionName{
			value: "telescope",
		},
	}
}

func (c ListAgentStatusRequestBodyExtensionName) Value() string {
	return c.value
}

func (c ListAgentStatusRequestBodyExtensionName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgentStatusRequestBodyExtensionName) UnmarshalJSON(b []byte) error {
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

type ListAgentStatusRequestBodyExtensionStatus struct {
	value string
}

type ListAgentStatusRequestBodyExtensionStatusEnum struct {
	NONE    ListAgentStatusRequestBodyExtensionStatus
	RUNNING ListAgentStatusRequestBodyExtensionStatus
	STOPPED ListAgentStatusRequestBodyExtensionStatus
	FAULT   ListAgentStatusRequestBodyExtensionStatus
	UNKNOWN ListAgentStatusRequestBodyExtensionStatus
}

func GetListAgentStatusRequestBodyExtensionStatusEnum() ListAgentStatusRequestBodyExtensionStatusEnum {
	return ListAgentStatusRequestBodyExtensionStatusEnum{
		NONE: ListAgentStatusRequestBodyExtensionStatus{
			value: "none",
		},
		RUNNING: ListAgentStatusRequestBodyExtensionStatus{
			value: "running",
		},
		STOPPED: ListAgentStatusRequestBodyExtensionStatus{
			value: "stopped",
		},
		FAULT: ListAgentStatusRequestBodyExtensionStatus{
			value: "fault",
		},
		UNKNOWN: ListAgentStatusRequestBodyExtensionStatus{
			value: "unknown",
		},
	}
}

func (c ListAgentStatusRequestBodyExtensionStatus) Value() string {
	return c.value
}

func (c ListAgentStatusRequestBodyExtensionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAgentStatusRequestBodyExtensionStatus) UnmarshalJSON(b []byte) error {
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
