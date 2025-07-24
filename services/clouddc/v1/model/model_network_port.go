package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NetworkPort 网络端口的详细信息，如型号、厂商等
type NetworkPort struct {

	// 网络端口的物理端口号
	PhysicalPortNumber *int32 `json:"physical_port_number,omitempty"`

	// 网络端口的物理连接状态
	LinkStatus *NetworkPortLinkStatus `json:"link_status,omitempty"`

	// 网络端口的网络地址
	AssociatedNetworkAddresses *string `json:"associated_network_addresses,omitempty"`

	// 网络端口的网络协议
	ActiveLinkTechnology *string `json:"active_link_technology,omitempty"`

	// 网络端口的网口类型
	PortType *string `json:"port_type,omitempty"`

	// 网络端口的最大速率
	PortMaxSpeed *string `json:"port_max_speed,omitempty"`

	// 网络端口的固件版本
	FirmwarePackageVersion *string `json:"firmware_package_version,omitempty"`

	// 网络端口的BDF
	Bdf *string `json:"bdf,omitempty"`

	// 协议
	AutoNeg *string `json:"auto_neg,omitempty"`
}

func (o NetworkPort) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkPort struct{}"
	}

	return strings.Join([]string{"NetworkPort", string(data)}, " ")
}

type NetworkPortLinkStatus struct {
	value string
}

type NetworkPortLinkStatusEnum struct {
	UP   NetworkPortLinkStatus
	DOWN NetworkPortLinkStatus
}

func GetNetworkPortLinkStatusEnum() NetworkPortLinkStatusEnum {
	return NetworkPortLinkStatusEnum{
		UP: NetworkPortLinkStatus{
			value: "Up",
		},
		DOWN: NetworkPortLinkStatus{
			value: "Down",
		},
	}
}

func (c NetworkPortLinkStatus) Value() string {
	return c.value
}

func (c NetworkPortLinkStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NetworkPortLinkStatus) UnmarshalJSON(b []byte) error {
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
