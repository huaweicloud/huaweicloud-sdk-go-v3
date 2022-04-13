package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 创建网络实例的详细信息。
type CreateNetworkInstance struct {
	// 网络实例的名字。

	Name *string `json:"name,omitempty"`
	// 网络实例的描述。

	Description *string `json:"description,omitempty"`
	// 添加到云连接网络实例的类型。|- VPC：虚拟私有云。 VGW：虚拟网关。 ER：虚拟路由器。

	Type CreateNetworkInstanceType `json:"type"`
	// 添加到云连接网络实例的ID，VPC或者VGW的ID。

	InstanceId string `json:"instance_id"`
	// VPC或者VGW的账户ID。

	InstanceDomainId *string `json:"instance_domain_id,omitempty"`
	// VPC或者VGW的项目ID。

	ProjectId string `json:"project_id"`
	// VPC或者VGW的RegionID。

	RegionId string `json:"region_id"`
	// 云连接实例ID。

	CloudConnectionId string `json:"cloud_connection_id"`
	// VPC或者VGW发布的网段路由列表，ER场景不需要此字段。

	Cidrs []string `json:"cidrs"`
}

func (o CreateNetworkInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNetworkInstance struct{}"
	}

	return strings.Join([]string{"CreateNetworkInstance", string(data)}, " ")
}

type CreateNetworkInstanceType struct {
	value string
}

type CreateNetworkInstanceTypeEnum struct {
	VPC CreateNetworkInstanceType
	VGW CreateNetworkInstanceType
	ER  CreateNetworkInstanceType
}

func GetCreateNetworkInstanceTypeEnum() CreateNetworkInstanceTypeEnum {
	return CreateNetworkInstanceTypeEnum{
		VPC: CreateNetworkInstanceType{
			value: "vpc",
		},
		VGW: CreateNetworkInstanceType{
			value: "vgw",
		},
		ER: CreateNetworkInstanceType{
			value: "er",
		},
	}
}

func (c CreateNetworkInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNetworkInstanceType) UnmarshalJSON(b []byte) error {
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
