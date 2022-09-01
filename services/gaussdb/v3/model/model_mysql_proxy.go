package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlProxy struct {

	// Proxy实例id。
	PoolId *string `json:"pool_id,omitempty" xml:"pool_id"`

	// Proxy实例开启状态。  取值范围：closed、open、frozen、opening、closing、enlarging、freezing和unfreezing。
	Status *string `json:"status,omitempty" xml:"status"`

	// Proxy读写分离地址。
	Address *string `json:"address,omitempty" xml:"address"`

	// Proxy端口信息。
	Port *int32 `json:"port,omitempty" xml:"port"`

	// Proxy实例状态。 取值范围：ACTIVE、BUILD、FAILED和DELETED。
	PoolStatus *string `json:"pool_status,omitempty" xml:"pool_status"`

	// 延时阈值，单位：秒。
	DelayThresholdInSeconds *int32 `json:"delay_threshold_in_seconds,omitempty" xml:"delay_threshold_in_seconds"`

	// Elb模式的虚拟ip信息。
	ElbVip *string `json:"elb_vip,omitempty" xml:"elb_vip"`

	// 弹性公网IP信息。
	Eip *string `json:"eip,omitempty" xml:"eip"`

	// Proxy实例规格的CPU数量。
	Vcpus *string `json:"vcpus,omitempty" xml:"vcpus"`

	// Proxy实例规格的内存数量。
	Ram *string `json:"ram,omitempty" xml:"ram"`

	// Proxy节点个数。
	NodeNum *int32 `json:"node_num,omitempty" xml:"node_num"`

	// Proxy主备模式，取值范围：Cluster。
	Mode *string `json:"mode,omitempty" xml:"mode"`

	// Proxy节点信息。
	Nodes *[]MysqlProxyNodes `json:"nodes,omitempty" xml:"nodes"`

	// Proxy规格信息。
	FlavorRef *string `json:"flavor_ref,omitempty" xml:"flavor_ref"`

	// Proxy实例名称。
	Name *string `json:"name,omitempty" xml:"name"`
}

func (o MysqlProxy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlProxy struct{}"
	}

	return strings.Join([]string{"MysqlProxy", string(data)}, " ")
}
