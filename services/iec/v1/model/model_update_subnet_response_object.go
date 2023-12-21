package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSubnetResponseObject 更新子网响应对象
type UpdateSubnetResponseObject struct {

	// 子网ID
	Id *string `json:"id,omitempty"`

	// 子网的状态  取值范围： - ACTIVE：表示子网已挂载到ROUTER上 - UNKNOWN：表示子网还未挂载到ROUTER上 - ERROR：表示子网状态故障
	Status *UpdateSubnetResponseObjectStatus `json:"status,omitempty"`

	// 是否开启IPv6
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// 对应IPv6子网（OpenStack Neutron接口）id，如果子网为IPv4子网，则不返回此参数。
	NeutronSubnetIdV6 *string `json:"neutron_subnet_id_v6,omitempty"`
}

func (o UpdateSubnetResponseObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubnetResponseObject struct{}"
	}

	return strings.Join([]string{"UpdateSubnetResponseObject", string(data)}, " ")
}

type UpdateSubnetResponseObjectStatus struct {
	value string
}

type UpdateSubnetResponseObjectStatusEnum struct {
	ACTIVE  UpdateSubnetResponseObjectStatus
	UNKNOWN UpdateSubnetResponseObjectStatus
	ERROR_  UpdateSubnetResponseObjectStatus
}

func GetUpdateSubnetResponseObjectStatusEnum() UpdateSubnetResponseObjectStatusEnum {
	return UpdateSubnetResponseObjectStatusEnum{
		ACTIVE: UpdateSubnetResponseObjectStatus{
			value: "ACTIVE",
		},
		UNKNOWN: UpdateSubnetResponseObjectStatus{
			value: "UNKNOWN",
		},
		ERROR_: UpdateSubnetResponseObjectStatus{
			value: "ERROR  ",
		},
	}
}

func (c UpdateSubnetResponseObjectStatus) Value() string {
	return c.value
}

func (c UpdateSubnetResponseObjectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSubnetResponseObjectStatus) UnmarshalJSON(b []byte) error {
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
