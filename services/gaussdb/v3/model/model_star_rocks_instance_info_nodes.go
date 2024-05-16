package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StarRocksInstanceInfoNodes struct {

	// 实例节点ID。
	Id *string `json:"id,omitempty"`

	// 实例节点名。
	Name *string `json:"name,omitempty"`

	// 实例节点类型。
	Type *string `json:"type,omitempty"`

	// 节点状态。  取值：  值为“creating”，表示节点正在创建。  值为“normal”，表示节点正常。  值为“abnormal”，表示节点异常。  值为“createfail”，表示节点创建失败。
	Status *string `json:"status,omitempty"`

	// 实例节点周期。
	Period *string `json:"period,omitempty"`

	Volume *StarRocksInstanceInfoVolume `json:"volume,omitempty"`

	// 实例节点cpu数量。
	Cpu *string `json:"cpu,omitempty"`

	// 实例节点内存大小（GB）。
	Mem *string `json:"mem,omitempty"`

	Datastore *StarRocksInstanceInfoDatastore `json:"datastore,omitempty"`

	// 节点动作。
	Actions *[]QueryAction `json:"actions,omitempty"`

	// 节点优先级。
	Priority *int32 `json:"priority,omitempty"`

	// 冻结标志。
	FrozenFlag *int32 `json:"frozen_flag,omitempty"`

	// 数据库端口号。默认3306。
	DbPort *int32 `json:"db_port,omitempty"`

	// 支付模式。
	PayModel *string `json:"pay_model,omitempty"`

	// 订单号。
	OrderId *string `json:"order_id,omitempty"`

	// 数据IP。
	TrafficIp *string `json:"traffic_ip,omitempty"`

	// 数据IPV6。
	TrafficIpv6 *string `json:"traffic_ipv6,omitempty"`

	// 可用区代码。
	AzCode *string `json:"az_code,omitempty"`

	// 可用区描述。
	AzDescription *string `json:"az_description,omitempty"`

	// 可用区类型。
	AzType *string `json:"az_type,omitempty"`

	// 实例所在区域。
	RegionCode *string `json:"region_code,omitempty"`

	// 节点创建时间。
	CreateAt *int64 `json:"create_at,omitempty"`

	// 节点更新时间。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// 节点规格ID。
	FlavorId *string `json:"flavor_id,omitempty"`

	// 节点规格码。
	FlavorRef *string `json:"flavor_ref,omitempty"`

	// IASS规格码。
	IassFlavorRef *string `json:"iass_flavor_ref,omitempty"`

	// 公网最大连接数。
	MaxConnections *string `json:"max_connections,omitempty"`

	// 虚拟私有云ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 参数更新是否需要重启。
	NeedRestart *bool `json:"need_restart,omitempty"`

	// 安全组。
	SgId *string `json:"sg_id,omitempty"`

	// 参数组信息。
	ParamGroup map[string]ParamGroup `json:"param_group,omitempty"`
}

func (o StarRocksInstanceInfoNodes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksInstanceInfoNodes struct{}"
	}

	return strings.Join([]string{"StarRocksInstanceInfoNodes", string(data)}, " ")
}
