package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EgressCidr CAE环境访问VPC配置。
type EgressCidr struct {

	// 目的网络Cidr。
	Cidr string `json:"cidr"`

	// 目的网络所属CAE环境VPC的路由表ID。
	RouteTableId string `json:"route_table_id"`

	// CAE环境访问VPC配置ID。
	Id *string `json:"id,omitempty"`
}

func (o EgressCidr) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EgressCidr struct{}"
	}

	return strings.Join([]string{"EgressCidr", string(data)}, " ")
}
