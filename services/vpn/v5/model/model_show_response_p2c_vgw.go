package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowResponseP2cVgw struct {

	// P2C VPN网关ID
	Id *string `json:"id,omitempty"`

	// P2C VPN网关名称
	Name *string `json:"name,omitempty"`

	// P2C VPN网关状态
	Status *string `json:"status,omitempty"`

	// P2C VPN网关所连接的VPC的ID
	VpcId *string `json:"vpc_id,omitempty"`

	// P2C VPN网关所使用的VPC子网ID
	ConnectSubnet *string `json:"connect_subnet,omitempty"`

	// P2C VPN网关的规格类型
	Flavor *string `json:"flavor,omitempty"`

	// 可用区列表
	AvailabilityZoneIds *[]string `json:"availability_zone_ids,omitempty"`

	Eip *ResponseEipInfo `json:"eip,omitempty"`

	// 配置的最大并发客户端连接数
	MaxConnectionNumber *int32 `json:"max_connection_number,omitempty"`

	// 当前建连的客户端连接数
	CurrentConnectionNumber *int32 `json:"current_connection_number,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签
	Tags *[]VpnResourceTag `json:"tags,omitempty"`

	// 订单Id
	OrderId *string `json:"order_id,omitempty"`

	// 冻结状态
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 冻结场景：0未冻结；1 冻结可删除；2冻结不可删除
	FrozenEffect *int32 `json:"frozen_effect,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o ShowResponseP2cVgw) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResponseP2cVgw struct{}"
	}

	return strings.Join([]string{"ShowResponseP2cVgw", string(data)}, " ")
}
