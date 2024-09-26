package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CloudConnectionRoute 云连接路由实例。
type CloudConnectionRoute struct {

	// 实例ID。
	Id string `json:"id"`

	// 云连接实例ID。
	CloudConnectionId string `json:"cloud_connection_id"`

	// 实例所属账号ID。
	DomainId string `json:"domain_id"`

	// 网络实例（VPC，VGW）的ID。
	InstanceId string `json:"instance_id"`

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`

	// RegionID。
	RegionId string `json:"region_id"`

	// 路由条目下一跳指向的网络实例的类型。 - VPC：虚拟私有云。 - VGW：虚拟网关。
	Type *CloudConnectionRouteType `json:"type,omitempty"`

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

func (c CloudConnectionRouteType) Value() string {
	return c.value
}

func (c CloudConnectionRouteType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudConnectionRouteType) UnmarshalJSON(b []byte) error {
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
