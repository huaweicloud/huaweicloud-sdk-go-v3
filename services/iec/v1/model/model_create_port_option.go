package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePortOption 创建端口参数对象。
type CreatePortOption struct {

	// 端口设备所属。  取值范围：目前只支持指定\"neutron:VIP_PORT\"，neutron:VIP_PORT表示创建的是VIP
	DeviceOwner CreatePortOptionDeviceOwner `json:"device_owner"`

	// 端口所属网络的ID。  约束：必须是存在的网络ID。
	NetworkId string `json:"network_id"`

	// 端口IP  约束：一个端口只支持一个fixed_ip，且不支持更新。
	FixedIps *[]FixedIp `json:"fixed_ips,omitempty"`
}

func (o CreatePortOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePortOption struct{}"
	}

	return strings.Join([]string{"CreatePortOption", string(data)}, " ")
}

type CreatePortOptionDeviceOwner struct {
	value string
}

type CreatePortOptionDeviceOwnerEnum struct {
	NEUTRONVIP_PORT CreatePortOptionDeviceOwner
}

func GetCreatePortOptionDeviceOwnerEnum() CreatePortOptionDeviceOwnerEnum {
	return CreatePortOptionDeviceOwnerEnum{
		NEUTRONVIP_PORT: CreatePortOptionDeviceOwner{
			value: "neutron:VIP_PORT",
		},
	}
}

func (c CreatePortOptionDeviceOwner) Value() string {
	return c.value
}

func (c CreatePortOptionDeviceOwner) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePortOptionDeviceOwner) UnmarshalJSON(b []byte) error {
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
