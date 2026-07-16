package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolSpecModelContainernetwork **参数解释**：k8s 容器网络。
type PoolSpecModelContainernetwork struct {

	// **参数解释**：容器网络模型。 **取值范围**：可选值如下： - overlay_l2：容器隧道网络，通过OVS（OpenVSwitch）为容器构建的overlay_l2网络。 - vpc-router：VPC网络，使用ipvlan和自定义VPC路由为容器构建的Underlay的l2网络。 - eni：云原生网络2.0，深度整合VPC原生ENI弹性网卡能力，采用VPC网段分配容器地址，支持ELB直通容器，享有高性能，创建CCE Turbo集群时指定。
	Mode *string `json:"mode,omitempty"`
}

func (o PoolSpecModelContainernetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolSpecModelContainernetwork struct{}"
	}

	return strings.Join([]string{"PoolSpecModelContainernetwork", string(data)}, " ")
}
