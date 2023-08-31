package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ListAgentStatusRequestBody struct {

	// 机器实例id列表
	InstanceIds []string `json:"instance_ids"`

	// uniagent运行状态，不传查所有状态,none无，running运行中，silent静默中，unknown故障
	UniagentStatus *ListAgentStatusRequestBodyUniagentStatus `json:"uniagent_status,omitempty"`

	// 插件名称，不传查所有插件，目前仅支持telescope
	ExtensionName *ListAgentStatusRequestBodyExtensionName `json:"extension_name,omitempty"`

	// 插件状态，不传查所有状态, none未安装，running运行中，stopped已停止，fault故障（进程异常），unknown故障（连接异常）
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
