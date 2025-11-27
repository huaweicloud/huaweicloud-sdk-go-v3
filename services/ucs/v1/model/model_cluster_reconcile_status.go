package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterReconcileStatus 针对单个特定集群的综合状态概览
type ClusterReconcileStatus struct {

	// 集群id
	ClusterID *string `json:"clusterID,omitempty"`

	// 集群中配置集合的总数
	ConfigSetTotalNum *int32 `json:"configSetTotalNum,omitempty"`

	// 健康状态的配置集合数量
	HealthStatusNum *int32 `json:"healthStatusNum,omitempty"`

	// 失败状态的配置集合数量
	FailedStatusNum *int32 `json:"failedStatusNum,omitempty"`

	// 未知状态的配置集合数量
	UnknownStatusNum *int32 `json:"unknownStatusNum,omitempty"`

	// 正在处理中的配置集合数量
	ProgressingStatusNum *int32 `json:"progressingStatusNum,omitempty"`

	// 与集群关联的Kubernetes资源数量
	K8sResourceNum *int32 `json:"k8sResourceNum,omitempty"`

	// 与集群关联的源代码仓库数量
	SourceRepoNum *int32 `json:"sourceRepoNum,omitempty"`
}

func (o ClusterReconcileStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterReconcileStatus struct{}"
	}

	return strings.Join([]string{"ClusterReconcileStatus", string(data)}, " ")
}
