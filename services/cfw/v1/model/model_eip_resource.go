package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EipResource EIP资源防护信息
type EipResource struct {

	// 弹性公网ID
	Id *string `json:"id,omitempty"`

	// 弹性公网IP
	PublicIp *string `json:"public_ip,omitempty"`

	// EIP防护状态，0表示防护中，1表示未防护
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

	// 防火墙名称
	FwInstanceName *string `json:"fw_instance_name,omitempty"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)。
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// Eip绑定的防火墙企业项目id
	FwEnterpriseProjectId *string `json:"fw_enterprise_project_id,omitempty"`

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用[查询防火墙实例接口](ListFirewallDetail.xml)，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。
	ObjectId *string `json:"object_id,omitempty"`

	// 标签列表
	Tags *string `json:"tags,omitempty"`

	// EIP所属用户
	DomainId *string `json:"domain_id,omitempty"`

	// 所属用户的名称
	Owner *string `json:"owner,omitempty"`

	// 防火墙所属用户
	FwDomainId *string `json:"fw_domain_id,omitempty"`
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
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
