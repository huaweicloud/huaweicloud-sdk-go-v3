/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

//
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
	data, _ := json.Marshal(o)
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

func (c CreateVpcRouteOptionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateVpcRouteOptionType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
