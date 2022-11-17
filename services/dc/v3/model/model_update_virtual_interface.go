package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 虚拟接口更新对象
type UpdateVirtualInterface struct {

	// 虚拟接口名字
	Name *string `json:"name,omitempty"`

	// 虚拟接口描述信息
	Description *string `json:"description,omitempty"`

	// 虚拟接口带宽配置
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 远端子网列表，记录租户侧的cidrs
	RemoteEpGroup *[]string `json:"remote_ep_group,omitempty"`

	// 用于公网专线,用户访问公网服务地址列表
	ServiceEpGroup *[]string `json:"service_ep_group,omitempty"`

	// 是否使能bfd功能：true或false
	EnableBfd *bool `json:"enable_bfd,omitempty"`

	// 是否使能nqa功能：true或false
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
