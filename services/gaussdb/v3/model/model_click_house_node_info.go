package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClickHouseNodeInfo 节点信息。
type ClickHouseNodeInfo struct {

	// 实例节点ID。
	Id string `json:"id"`

	// 实例节点名。
	Name string `json:"name"`

	// 实例节点类型。 取值范围： - master：主节点 - slave：备节点
	Type string `json:"type"`

	// 实例节点状态。
	Status string `json:"status"`

	// 实例节点周期。
	Period *string `json:"period,omitempty"`

	Volume *ClickHouseNodeInfoVolume `json:"volume"`

	// 实例节点cpu数量。
	Cpu string `json:"cpu"`

	// 实例节点内存大小（GB）。
	Mem string `json:"mem"`

	Datastore *ClickHouseNodeInfoDatastore `json:"datastore"`

	// 节点优先级。
	Priority int32 `json:"priority"`

	// 冻结标志。 取值范围： - 0：不冻结 - 1：冻结
	FrozenFlag int32 `json:"frozen_flag"`

	// 数据库端口号。取值范围：0~65535。
	DbPort int32 `json:"db_port"`

	// 支付模式。 取值范围： - 0：按需计费 - 1：包周期
	PayModel string `json:"pay_model"`

	// 包周期订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 节点数据ip。
	TrafficIp string `json:"traffic_ip"`

	// 节点数据ipv6。
	TrafficIpv6 *string `json:"traffic_ipv6,omitempty"`

	// 节点数据vip。
	TrafficVip *string `json:"traffic_vip,omitempty"`

	// 节点数据vipV6。
	TrafficVipv6 *string `json:"traffic_vipv6,omitempty"`

	// 可用区。
	AzCode string `json:"az_code"`

	// 可用区描述。
	AzDescription string `json:"az_description"`

	// 可用区类型。
	AzType string `json:"az_type"`

	// 节点所在区。
	Region string `json:"region"`

	// 节点创建时间。
	CreateAt int64 `json:"create_at"`

	// 节点更新时间。
	UpdateAt int64 `json:"update_at"`

	// 节点规格ID。
	FlavorId string `json:"flavor_id"`

	// 节点规格码。
	FlavorRef string `json:"flavor_ref"`

	// IASS规格码。
	IassFlavorRef string `json:"iass_flavor_ref"`

	// 公网最大连接数。
	MaxConnections string `json:"max_connections"`

	// 虚拟私有云ID。
	VpcId string `json:"vpc_id"`

	// 子网ID。
	SubnetId string `json:"subnet_id"`

	// 参数更新是否需要重启。
	NeedRestart bool `json:"need_restart"`

	// 安全组
	SgId string `json:"sg_id"`
}

func (o ClickHouseNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClickHouseNodeInfo struct{}"
	}

	return strings.Join([]string{"ClickHouseNodeInfo", string(data)}, " ")
}
