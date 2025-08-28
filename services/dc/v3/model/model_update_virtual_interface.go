package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateVirtualInterface 虚拟接口更新对象
type UpdateVirtualInterface struct {

	// 虚拟接口名字
	Name *string `json:"name,omitempty"`

	// 虚拟接口描述信息
	Description *string `json:"description,omitempty"`

	// 虚拟接口带宽配置
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 虚拟接口的优先级，支持两种优先级状态normal和low。 接口优先级相同时表示负载关系，接口优先级不同时表示主备关系，出云流量优先转到优先级更高的normal接口。 目前仅BGP模式接口支持。
	Priority *UpdateVirtualInterfacePriority `json:"priority,omitempty"`

	// 远端子网列表，记录租户侧的cidrs
	RemoteEpGroup *[]string `json:"remote_ep_group,omitempty"`

	// 用于公网专线,用户访问公网服务地址列表。[（预留字段，暂不支持）](tag:dt)
	ServiceEpGroup *[]string `json:"service_ep_group,omitempty"`

	// 是否使能bfd功能：true或false。[（预留字段，暂不支持）](tag:dt)
	EnableBfd *bool `json:"enable_bfd,omitempty"`

	// 是否使能nqa功能：true或false。[（预留字段，暂不支持）](tag:dt)
	EnableNqa *bool `json:"enable_nqa,omitempty"`

	// 对其他租户创建的虚拟接口进行确认,可以是ACCEPTED和REJECTED
	Status *UpdateVirtualInterfaceStatus `json:"status,omitempty"`
}

func (o UpdateVirtualInterface) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVirtualInterface struct{}"
	}

	return strings.Join([]string{"UpdateVirtualInterface", string(data)}, " ")
}

type UpdateVirtualInterfacePriority struct {
	value string
}

type UpdateVirtualInterfacePriorityEnum struct {
	NORMAL UpdateVirtualInterfacePriority
	LOW    UpdateVirtualInterfacePriority
}

func GetUpdateVirtualInterfacePriorityEnum() UpdateVirtualInterfacePriorityEnum {
	return UpdateVirtualInterfacePriorityEnum{
		NORMAL: UpdateVirtualInterfacePriority{
			value: "normal",
		},
		LOW: UpdateVirtualInterfacePriority{
			value: "low",
		},
	}
}

func (c UpdateVirtualInterfacePriority) Value() string {
	return c.value
}

func (c UpdateVirtualInterfacePriority) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateVirtualInterfacePriority) UnmarshalJSON(b []byte) error {
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

type UpdateVirtualInterfaceStatus struct {
	value string
}

type UpdateVirtualInterfaceStatusEnum struct {
	ACCEPTED UpdateVirtualInterfaceStatus
	REJECTED UpdateVirtualInterfaceStatus
}

func GetUpdateVirtualInterfaceStatusEnum() UpdateVirtualInterfaceStatusEnum {
	return UpdateVirtualInterfaceStatusEnum{
		ACCEPTED: UpdateVirtualInterfaceStatus{
			value: "ACCEPTED",
		},
		REJECTED: UpdateVirtualInterfaceStatus{
			value: "REJECTED",
		},
	}
}

func (c UpdateVirtualInterfaceStatus) Value() string {
	return c.value
}

func (c UpdateVirtualInterfaceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateVirtualInterfaceStatus) UnmarshalJSON(b []byte) error {
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
