package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// NetworkInstance 网络实例。
type NetworkInstance struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 实例所属账号ID。
	DomainId string `json:"domain_id"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 云连接实例ID。
	CloudConnectionId string `json:"cloud_connection_id"`

	// 网络实例（VPC，VGW）的ID。
	InstanceId string `json:"instance_id"`

	// 网络实例（VPC，VGW）所属账号ID。
	InstanceDomainId *string `json:"instance_domain_id,omitempty"`

	// RegionID。
	RegionId string `json:"region_id"`

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`

	// 网络实例的状态。 - ACTIVE：处理成功。 - PENDING：处理中。 - ERROR：处理失败。
	Status *NetworkInstanceStatus `json:"status,omitempty"`

	// 网络实例的类型。 - vpc：虚拟私有云。 - vgw：虚拟网关。
	Type *NetworkInstanceType `json:"type,omitempty"`

	// 网络实例发布的网段路由列表。
	Cidrs *[]string `json:"cidrs,omitempty"`
}

func (o NetworkInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkInstance struct{}"
	}

	return strings.Join([]string{"NetworkInstance", string(data)}, " ")
}

type NetworkInstanceStatus struct {
	value string
}

type NetworkInstanceStatusEnum struct {
	ACTIVE NetworkInstanceStatus
}

func GetNetworkInstanceStatusEnum() NetworkInstanceStatusEnum {
	return NetworkInstanceStatusEnum{
		ACTIVE: NetworkInstanceStatus{
			value: "ACTIVE",
		},
	}
}

func (c NetworkInstanceStatus) Value() string {
	return c.value
}

func (c NetworkInstanceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NetworkInstanceStatus) UnmarshalJSON(b []byte) error {
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

type NetworkInstanceType struct {
	value string
}

type NetworkInstanceTypeEnum struct {
	VPC NetworkInstanceType
	VGW NetworkInstanceType
}

func GetNetworkInstanceTypeEnum() NetworkInstanceTypeEnum {
	return NetworkInstanceTypeEnum{
		VPC: NetworkInstanceType{
			value: "vpc",
		},
		VGW: NetworkInstanceType{
			value: "vgw",
		},
	}
}

func (c NetworkInstanceType) Value() string {
	return c.value
}

func (c NetworkInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NetworkInstanceType) UnmarshalJSON(b []byte) error {
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
