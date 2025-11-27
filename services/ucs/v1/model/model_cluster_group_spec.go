package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusterGroupSpec struct {

	// 权限策略关联的命名空间列表
	RuleNamespaces *[]string `json:"ruleNamespaces,omitempty"`

	// 舰队启用联邦ID
	FederationId *string `json:"federationId,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 舰队对应联邦的DNS后缀，开启联邦后可见
	DnsSuffix *[]string `json:"dnsSuffix,omitempty"`

	// 联邦到期时间戳
	FederationExpirationTimestamp *string `json:"federationExpirationTimestamp,omitempty"`

	// 策略管理id
	PolicyId *string `json:"policyId,omitempty"`

	// 舰队启用联邦的版本
	FederationVersion *string `json:"federationVersion,omitempty"`

	// 集群联邦连接网关节点列表，仅当舰队开启集群联邦时有值
	ConnectGatewayEndpoints *[]ConnectEndpoint `json:"connectGatewayEndpoints,omitempty"`
}

func (o ClusterGroupSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterGroupSpec struct{}"
	}

	return strings.Join([]string{"ClusterGroupSpec", string(data)}, " ")
}
