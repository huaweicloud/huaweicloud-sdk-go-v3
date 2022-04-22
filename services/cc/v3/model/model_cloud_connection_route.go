package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 云连接路由实例。
type CloudConnectionRoute struct {

	// 云连接实例路由的ID。
	Id *string `json:"id,omitempty"`

	// 云连接实例的ID。
	CloudConnectionId *string `json:"cloud_connection_id,omitempty"`

	// 帐号ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 网络实例的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 路由条目下一跳指向的网络实例的ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 路由条目下一跳指向的网络实例的类型。 |- VPC：虚拟私有云。 VGW：虚拟网关。
	Type *CloudConnectionRouteType `json:"type,omitempty"`

	// Region的ID。
	RegionId *string `json:"region_id,omitempty"`

	// 目的地址。
	Destination *string `json:"destination,omitempty"`
}

func (o CloudConnectionRoute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudConnectionRoute struct{}"
	}

	return strings.Join([]string{"CloudConnectionRoute", string(data)}, " ")
}

type CloudConnectionRouteType struct {
	value string
}

type CloudConnectionRouteTypeEnum struct {
	VPC CloudConnectionRouteType
	VGW CloudConnectionRouteType
}

func GetCloudConnectionRouteTypeEnum() CloudConnectionRouteTypeEnum {
	return CloudConnectionRouteTypeEnum{
		VPC: CloudConnectionRouteType{
			value: "vpc",
		},
		VGW: CloudConnectionRouteType{
			value: "vgw",
		},
	}
}

func (c CloudConnectionRouteType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudConnectionRouteType) UnmarshalJSON(b []byte) error {
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
