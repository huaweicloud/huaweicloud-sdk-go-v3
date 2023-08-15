package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateVpcRouteOption
type CreateVpcRouteOption struct {

	// 路由目的地址CIDR，如192.168.200.0/24。
	Destination string `json:"destination"`

	// 功能说明：路由下一跳  取值范围：如果type为peering类型，则nexthop为peering的ID
	Nexthop string `json:"nexthop"`

	// 功能说明：路由类型  取值范围：peering
	Type CreateVpcRouteOptionType `json:"type"`

	// 请求添加路由的VPC ID
	VpcId string `json:"vpc_id"`
}

func (o CreateVpcRouteOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcRouteOption struct{}"
	}

	return strings.Join([]string{"CreateVpcRouteOption", string(data)}, " ")
}

type CreateVpcRouteOptionType struct {
	value string
}

type CreateVpcRouteOptionTypeEnum struct {
	PEERING CreateVpcRouteOptionType
}

func GetCreateVpcRouteOptionTypeEnum() CreateVpcRouteOptionTypeEnum {
	return CreateVpcRouteOptionTypeEnum{
		PEERING: CreateVpcRouteOptionType{
			value: "peering",
		},
	}
}

func (c CreateVpcRouteOptionType) Value() string {
	return c.value
}

func (c CreateVpcRouteOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVpcRouteOptionType) UnmarshalJSON(b []byte) error {
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
