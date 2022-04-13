package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/sdktime"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

// 网络实例。
type NetworkInstance struct {
	// 网络实例的ID。

	Id *string `json:"id,omitempty"`
	// 网络实例的名字。

	Name *string `json:"name,omitempty"`
	// 网络实例的描述。

	Description *string `json:"description,omitempty"`
	// 帐号ID。

	DomainId *string `json:"domain_id,omitempty"`
	// 网络实例的状态。ACTIVE：表示状态可用。

	Status *NetworkInstanceStatus `json:"status,omitempty"`
	// 网络实例的创建时间。

	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
	// 网络实例的更新时间。

	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
	// 网络实例的类型。|- VPC：虚拟私有云。 VGW：虚拟网关。 ER：虚拟路由器。

	Type *NetworkInstanceType `json:"type,omitempty"`
	// 云连接实例ID。

	CloudConnectionId *string `json:"cloud_connection_id,omitempty"`
	// VPC或者VGW的ID。

	InstanceId *string `json:"instance_id,omitempty"`
	// VPC或者VGW所属账户ID。

	InstanceDomainId *string `json:"instance_domain_id,omitempty"`
	// VPC或者VGW所在Region的ID。

	RegionId *string `json:"region_id,omitempty"`
	// VPC或者VGW所在租户的项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// VPC或者VGW发布的网段路由列表。

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

func (c NetworkInstanceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NetworkInstanceStatus) UnmarshalJSON(b []byte) error {
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

type NetworkInstanceType struct {
	value string
}

type NetworkInstanceTypeEnum struct {
	VPC NetworkInstanceType
	VGW NetworkInstanceType
	ER  NetworkInstanceType
}

func GetNetworkInstanceTypeEnum() NetworkInstanceTypeEnum {
	return NetworkInstanceTypeEnum{
		VPC: NetworkInstanceType{
			value: "vpc",
		},
		VGW: NetworkInstanceType{
			value: "vgw",
		},
		ER: NetworkInstanceType{
			value: "er",
		},
	}
}

func (c NetworkInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NetworkInstanceType) UnmarshalJSON(b []byte) error {
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
