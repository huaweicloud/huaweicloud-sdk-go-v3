package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建网络实例的详细信息。
type CreateNetworkInstance struct {

	// 网络实例的名字。只能由中文、英文字母、数字、下划线、中划线、点组成。
	Name *string `json:"name,omitempty"`

	// 网络实例的描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 添加到云连接网络实例的类型，有效值： - vpc：虚拟私有云。 - vgw：虚拟网关。
	Type CreateNetworkInstanceType `json:"type"`

	// 添加到云连接网络实例的ID，VPC或者VGW的ID。
	InstanceId string `json:"instance_id"`

	// 网络实例的账户ID。跨账号加载必填；同账号下资源加载不填。
	InstanceDomainId *string `json:"instance_domain_id,omitempty"`

	// 网络实例的项目ID。
	ProjectId string `json:"project_id"`

	// 网络实例的RegionID。
	RegionId string `json:"region_id"`

	// 云连接实例ID。
	CloudConnectionId string `json:"cloud_connection_id"`

	// 网络实例发布的网段路由列表。
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
}

func GetCreateNetworkInstanceTypeEnum() CreateNetworkInstanceTypeEnum {
	return CreateNetworkInstanceTypeEnum{
		VPC: CreateNetworkInstanceType{
			value: "vpc",
		},
		VGW: CreateNetworkInstanceType{
			value: "vgw",
		},
	}
}

func (c CreateNetworkInstanceType) Value() string {
	return c.value
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
