package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlProxyV3 struct {

	// Proxy实例id。
	PoolId *string `json:"pool_id,omitempty"`

	// Proxy实例开启状态。  取值: 值为“ACTIVE”，表示数据库代理正常； “FAILED”，表示数据库代理创建失败； “DELETED”，表示数据库代理已删除； “ABNORMAL”，表示数据库代理异常； “ENABLING PROXY”，表示数据库代理正在开启； “DISABLING PROXY”，表示数据库代理正在关闭； “ADDING PROXY NODE”，表示数据库代理正在扩容； “DELETING READ REPLICAS FROM PROXY”，表示数据库代理正在移除只读节点； “ADDING READ REPLICAS TO PROXY”，表示数据库代理正在添加只读节点； “CHANGING WEIGHTS”，表示数据库代理正在修改只读节点权重。
	Status *string `json:"status,omitempty"`

	// Proxy读写分离地址。
	Address *string `json:"address,omitempty"`

	// Proxy端口信息。
	Port *int32 `json:"port,omitempty"`

	// Proxy实例状态。 取值范围：ACTIVE、ABNORMAL、FAILED和DELETED。
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
}

func (o MysqlProxyV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlProxyV3 struct{}"
	}

	return strings.Join([]string{"MysqlProxyV3", string(data)}, " ")
}
