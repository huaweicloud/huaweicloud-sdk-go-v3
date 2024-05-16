package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HtapInstanceListInstances struct {

	// HTAP实例ID，严格匹配UUID规则。
	Id *string `json:"id,omitempty"`

	// HTAP实例名。
	Name *string `json:"name,omitempty"`

	// HTAP数据库引擎名。
	EngineName *string `json:"engine_name,omitempty"`

	// HTAP数据库引擎版本。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 租户在某一region下的project ID
	ProjectId *string `json:"project_id,omitempty"`

	InstanceState *HtapInstanceListInstanceState `json:"instance_state,omitempty"`

	// HTAP实例创建时间。
	CreateAt *int64 `json:"create_at,omitempty"`

	// HTAP实例是否冻结。
	IsFrozen *bool `json:"is_frozen,omitempty"`

	// HTAP实例部署模式。
	HaMode *string `json:"ha_mode,omitempty"`

	// 计费模式。当前仅支持按需计费。 0：按需计费 1：包周期
	PayModel *string `json:"pay_model,omitempty"`

	// 包周期计费订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 包周期计费备用订单ID。
	AlterOrderId *string `json:"alter_order_id,omitempty"`

	// 读写内网地址。
	DataVip *string `json:"data_vip,omitempty"`

	// 可读节点信息
	ReadableNodeInfos *[]ReadableNodeInfos `json:"readable_node_infos,omitempty"`

	// 代理IP。
	ProxyIps *[]string `json:"proxy_ips,omitempty"`

	// 读写内网地址IPV6。
	DataVipV6 *string `json:"data_vip_v6,omitempty"`

	// 数据库访问端口。
	Port *int32 `json:"port,omitempty"`

	// 可用区信息。
	AvailableZones *[]HtapInstanceListAvailableZones `json:"available_zones,omitempty"`

	// 实例动作。
	CurrentActions *[]QueryAction `json:"current_actions,omitempty"`

	// 存储类型。
	VolumeType *string `json:"volume_type,omitempty"`

	// 服务器类型。
	ServerType *string `json:"server_type,omitempty"`

	// 企业项目ID。如果帐户开通企业项目服务则该参数必选，未开启该参数不可选。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 专属资源池ID，只有开通专属资源池后才支持此参数。
	DedicatedResourceId *string `json:"dedicated_resource_id,omitempty"`

	Network *HtapInstanceListNetwork `json:"network,omitempty"`

	// ClickHouse主节点ID。
	ChMasterNodeId *string `json:"ch_master_node_id,omitempty"`

	// 节点个数。
	NodeNum *int32 `json:"node_num,omitempty"`
}

func (o HtapInstanceListInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapInstanceListInstances struct{}"
	}

	return strings.Join([]string{"HtapInstanceListInstances", string(data)}, " ")
}
