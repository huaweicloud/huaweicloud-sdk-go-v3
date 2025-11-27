package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceBindingSpec struct {

	// 要传播的Kubernetes资源引用
	Resource *interface{} `json:"resource,omitempty"`

	// 是否自动传播相关资源
	PropagateDeps *bool `json:"propagateDeps,omitempty"`

	// 定义每个副本的需求
	ReplicaRequirements *interface{} `json:"replicaRequirements,omitempty"`

	// 要创建的副本数量
	Replicas *int32 `json:"replicas,omitempty"`

	Placement *Placement `json:"placement,omitempty"`

	// 目标成员集群列表
	Clusters *[]TargetCluster `json:"clusters,omitempty"`
}

func (o ResourceBindingSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceBindingSpec struct{}"
	}

	return strings.Join([]string{"ResourceBindingSpec", string(data)}, " ")
}
