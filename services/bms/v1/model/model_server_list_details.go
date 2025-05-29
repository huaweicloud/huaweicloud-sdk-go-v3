package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ServerListDetails server字段数据结构说明
type ServerListDetails struct {

	// 裸金属服务器ID，格式为UUID
	Id string `json:"id"`

	// 创建裸金属服务器的用户ID，格式为UUID。
	UserId *string `json:"user_id,omitempty"`

	// 裸金属服务器名称
	Name string `json:"name"`

	// 裸金属服务器创建时间。时间戳格式为ISO 8601：YYYY-MM-DDTHH:MM:SSZ，例如：2019-05-22T03:30:52Z
	Created *string `json:"created,omitempty"`

	// 裸金属服务器更新时间。时间戳格式为ISO 8601：YYYY-MM-DDTHH:MM:SSZ，例如：2019-05-22T04:30:52Z
	Updated *string `json:"updated,omitempty"`

	// 裸金属服务器所属租户ID，格式为UUID。该参数和project_id表示相同的概念。
	TenantId string `json:"tenant_id"`

	Flavor *FlavorDetailInfos `json:"flavor,omitempty"`

	// 裸金属服务器当前状态信息。  取值范围：  ACTIVE：运行中/正在关机/删除中 BUILD：创建中 ERROR：故障 HARD_REBOOT：强制重启中 REBOOT：重启中 DELETED：实例已被正常删除 SHUTOFF：关机/正在开机/删除中/重建中/重装操作系统中/重装操作系统失败/冻结
	Status ServerListDetailsStatus `json:"status"`

	// 扩展属性，裸金属服务器当前的任务状态。例如：rebooting：重启中reboot_started：普通重启reboot_started_hard：强制重启powering-off：关机中powering-on：开机中rebuilding：重建中scheduling：调度中deleting：删除中
	TaskState *ServerListDetailsTaskState `json:"task_state,omitempty"`

	// 扩展属性，裸金属服务器的稳定状态。例如：active：运行中shutoff：关机reboot：重启
	VmState *ServerListDetailsVmState `json:"vm_state,omitempty"`

	// 扩展属性，裸金属服务器所在可用分区名称。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	Fault *Fault `json:"fault,omitempty"`

	// 裸机是否在回收站中
	InRecycleBin *bool `json:"in_recycle_bin,omitempty"`
}

func (o ServerListDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerListDetails struct{}"
	}

	return strings.Join([]string{"ServerListDetails", string(data)}, " ")
}

type ServerListDetailsStatus struct {
	value string
}

type ServerListDetailsStatusEnum struct {
	ACTIVE      ServerListDetailsStatus
	BUILD       ServerListDetailsStatus
	ERROR       ServerListDetailsStatus
	HARD_REBOOT ServerListDetailsStatus
	REBOOT      ServerListDetailsStatus
	DELETED     ServerListDetailsStatus
	SHUTOFF     ServerListDetailsStatus
}

func GetServerListDetailsStatusEnum() ServerListDetailsStatusEnum {
	return ServerListDetailsStatusEnum{
		ACTIVE: ServerListDetailsStatus{
			value: "ACTIVE",
		},
		BUILD: ServerListDetailsStatus{
			value: "BUILD",
		},
		ERROR: ServerListDetailsStatus{
			value: "ERROR",
		},
		HARD_REBOOT: ServerListDetailsStatus{
			value: "HARD_REBOOT",
		},
		REBOOT: ServerListDetailsStatus{
			value: "REBOOT",
		},
		DELETED: ServerListDetailsStatus{
			value: "DELETED",
		},
		SHUTOFF: ServerListDetailsStatus{
			value: "SHUTOFF",
		},
	}
}

func (c ServerListDetailsStatus) Value() string {
	return c.value
}

func (c ServerListDetailsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerListDetailsStatus) UnmarshalJSON(b []byte) error {
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

type ServerListDetailsTaskState struct {
	value string
}

type ServerListDetailsTaskStateEnum struct {
	REBOOTING           ServerListDetailsTaskState
	REBOOT_STARTED      ServerListDetailsTaskState
	REBOOT_STARTED_HARD ServerListDetailsTaskState
	POWERING_OFF        ServerListDetailsTaskState
	POWERING_ON         ServerListDetailsTaskState
	REBUILDING          ServerListDetailsTaskState
	SCHEDULING          ServerListDetailsTaskState
	DELETING            ServerListDetailsTaskState
}

func GetServerListDetailsTaskStateEnum() ServerListDetailsTaskStateEnum {
	return ServerListDetailsTaskStateEnum{
		REBOOTING: ServerListDetailsTaskState{
			value: "rebooting",
		},
		REBOOT_STARTED: ServerListDetailsTaskState{
			value: "reboot_started",
		},
		REBOOT_STARTED_HARD: ServerListDetailsTaskState{
			value: "reboot_started_hard",
		},
		POWERING_OFF: ServerListDetailsTaskState{
			value: "powering-off",
		},
		POWERING_ON: ServerListDetailsTaskState{
			value: "powering-on",
		},
		REBUILDING: ServerListDetailsTaskState{
			value: "rebuilding",
		},
		SCHEDULING: ServerListDetailsTaskState{
			value: "scheduling",
		},
		DELETING: ServerListDetailsTaskState{
			value: "deleting",
		},
	}
}

func (c ServerListDetailsTaskState) Value() string {
	return c.value
}

func (c ServerListDetailsTaskState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerListDetailsTaskState) UnmarshalJSON(b []byte) error {
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

type ServerListDetailsVmState struct {
	value string
}

type ServerListDetailsVmStateEnum struct {
	ACTIVE  ServerListDetailsVmState
	SHUTOFF ServerListDetailsVmState
	REBOOT  ServerListDetailsVmState
}

func GetServerListDetailsVmStateEnum() ServerListDetailsVmStateEnum {
	return ServerListDetailsVmStateEnum{
		ACTIVE: ServerListDetailsVmState{
			value: "active",
		},
		SHUTOFF: ServerListDetailsVmState{
			value: "shutoff",
		},
		REBOOT: ServerListDetailsVmState{
			value: "reboot",
		},
	}
}

func (c ServerListDetailsVmState) Value() string {
	return c.value
}

func (c ServerListDetailsVmState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerListDetailsVmState) UnmarshalJSON(b []byte) error {
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
