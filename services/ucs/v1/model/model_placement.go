package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Placement struct {
	ClusterAffinity *ClusterAffinity `json:"clusterAffinity,omitempty"`

	// 集群容忍规则
	ClusterTolerations *[]Toleration `json:"clusterTolerations,omitempty"`

	// 定义如何在成员集群间分配副本
	ReplicaScheduling *interface{} `json:"replicaScheduling,omitempty"`
}

func (o Placement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Placement struct{}"
	}

	return strings.Join([]string{"Placement", string(data)}, " ")
}
