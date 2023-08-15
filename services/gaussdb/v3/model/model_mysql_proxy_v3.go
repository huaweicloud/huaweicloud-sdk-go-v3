package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlProxyV3 struct {

	// Proxy实例ID。
	PoolId *string `json:"pool_id,omitempty"`

	// Proxy实例开启状态。  取值: - “ACTIVE”，表示数据库代理正常； - “FAILED”，表示数据库代理创建失败； - “DELETED”，表示数据库代理已删除； - “ABNORMAL”，表示数据库代理异常； - “ENABLING PROXY”，表示数据库代理正在开启； - “DISABLING PROXY”，表示数据库代理正在关闭； - “ADDING PROXY NODE”，表示数据库代理正在扩容； - “DELETING READ REPLICAS FROM PROXY”，表示数据库代理正在移除只读节点； - “ADDING READ REPLICAS TO PROXY”，表示数据库代理正在添加只读节点； - “CHANGING WEIGHTS”，表示数据库代理正在修改只读节点权重。
	Status *string `json:"status,omitempty"`

	// Proxy读写分离地址。
	Address *string `json:"address,omitempty"`

	// Proxy端口信息。
	Port *int32 `json:"port,omitempty"`

	// Proxy实例状态。  取值范围： - ACTIVE，表示数据库代理正常 - ABNORMAL，表示数据库代理异常 - FAILED，表示数据库代理创建失败 - DELETED，表示数据库代理已删除
	PoolStatus *string `json:"pool_status,omitempty"`

	// 延时阈值，单位：秒。
	DelayThresholdInSeconds *int32 `json:"delay_threshold_in_seconds,omitempty"`

	// Elb模式的虚拟ip信息。
	ElbVip *string `json:"elb_vip,omitempty"`

	// 弹性公网IP信息。
	Eip *string `json:"eip,omitempty"`

	// Proxy实例规格的CPU数量。
	Vcpus *string `json:"vcpus,omitempty"`

	// Proxy实例规格的内存数量。
	Ram *string `json:"ram,omitempty"`

	// Proxy节点个数。
	NodeNum *int32 `json:"node_num,omitempty"`

	// Proxy主备模式，取值范围：Cluster。
	Mode *string `json:"mode,omitempty"`

	// Proxy节点信息。
	Nodes *[]MysqlProxyNodes `json:"nodes,omitempty"`

	// Proxy规格信息。
	FlavorRef *string `json:"flavor_ref,omitempty"`

	// Proxy实例名称。
	Name *string `json:"name,omitempty"`

	// Proxy事务拆分开关状态【ON/OFF】。
	TransactionSplit *string `json:"transaction_split,omitempty"`

	// 连接池类型。  取值范围: - CLOSED: 关闭连接池。 - SESSION: 开启会话级连接池。
	ConnectionPoolType *string `json:"connection_pool_type,omitempty"`

	// 数据库代理版本是否支持会话级连接池。  取值范围: - true: 支持。 - false: 不支持。
	SwitchConnectionPoolTypeEnabled *bool `json:"switch_connection_pool_type_enabled,omitempty"`

	// 数据库代理路由模式，默认为权重负载模式。  取值范围: - 0，表示权重负载模式; - 1，表示负载均衡模式（数据库主节点不接受读请求）； - 2，表示负载均衡模式（数据库主节点接受读请求）。
	RouteMode *int32 `json:"route_mode,omitempty"`

	// 数据库代理版本是否支持负载均衡模式。  取值范围: - true 支持; - false 不支持。
	BalanceRouteModeEnabled *bool `json:"balance_route_mode_enabled,omitempty"`

	// 一致性模式。默认值为空，此时以会话一致性参数session_consistence为准。 - session: 会话一致性 - global: 全局一致性 - eventual: 最终一致性
	ConsistenceMode *string `json:"consistence_mode,omitempty"`
}

func (o MysqlProxyV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlProxyV3 struct{}"
	}

	return strings.Join([]string{"MysqlProxyV3", string(data)}, " ")
}
