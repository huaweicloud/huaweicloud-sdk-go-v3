package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EIP资源防护信息
type EipResource struct {

	// 弹性公网ID
	Id *string `json:"id,omitempty"`

	// 弹性公网IP
	PublicIp *string `json:"public_ip,omitempty"`

	// EIP防护状态
	Status *EipResourceStatus `json:"status,omitempty"`

	// 弹性公网IP,IPV6
	PublicIpv6 *string `json:"public_ipv6,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 设备id
	DeviceId *string `json:"device_id,omitempty"`

	// 设备名称
	DeviceName *string `json:"device_name,omitempty"`

	// 设备拥有者
	DeviceOwner *string `json:"device_owner,omitempty"`

	// 关联实例类型
	AssociateInstanceType *string `json:"associate_instance_type,omitempty"`
}

func (o EipResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipResource struct{}"
	}

	return strings.Join([]string{"EipResource", string(data)}, " ")
}

type EipResourceStatus struct {
	value int32
}

type EipResourceStatusEnum struct {
	E_0 EipResourceStatus
	E_1 EipResourceStatus
}

func GetEipResourceStatusEnum() EipResourceStatusEnum {
	return EipResourceStatusEnum{
		E_0: EipResourceStatus{
			value: 0,
		}, E_1: EipResourceStatus{
			value: 1,
		},
	}
}

func (c EipResourceStatus) Value() int32 {
	return c.value
}

func (c EipResourceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EipResourceStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
