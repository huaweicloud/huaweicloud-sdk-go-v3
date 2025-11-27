package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClustersResponseInfo struct {

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群类型:   - cce：cce集群   - k8s：k8s集群    - ali：阿里云集群   - tencent：腾讯云集群   - azure：微软云集群   - aws：亚马逊集群   - self_built_hw：华为云自建集群   - self_built_idc：IDC自建集群
	ClusterType *string `json:"cluster_type,omitempty"`

	// 集群版本
	Version *string `json:"version,omitempty"`

	// 网络模型:   - overlay_l2：容器隧道网络   - vpc-router：VPC网络   - eni：云原生网络2.0   - native-network：K8S原生网络
	Mode *string `json:"mode,omitempty"`

	// 命名空间数
	NamespaceNum *int32 `json:"namespace_num,omitempty"`

	// 策略数量
	PolicyNum *int32 `json:"policy_num,omitempty"`

	// **参数解释**: 防护状态 **取值范围**: - true: 防护中 - false: 未防护
	ProtectionStatus *bool `json:"protection_status,omitempty"`
}

func (o ClustersResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClustersResponseInfo struct{}"
	}

	return strings.Join([]string{"ClustersResponseInfo", string(data)}, " ")
}
