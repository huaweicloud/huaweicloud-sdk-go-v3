package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterDisasterRecovery 集群容灾恢复详情
type ClusterDisasterRecovery struct {
	PrimaryCluster *ClusterRecoveryProgress `json:"primary_cluster,omitempty"`

	StandbyCluster *ClusterRecoveryProgress `json:"standby_cluster,omitempty"`

	// latest_barrier_time
	LatestBarrierTime *string `json:"latest_barrier_time,omitempty"`

	// last_recovery_spend
	LastRecoverySpend *int64 `json:"last_recovery_spend,omitempty"`

	// recovery_point_object
	RecoveryPointObject *int64 `json:"recovery_point_object,omitempty"`

	// recovery_time_object
	RecoveryTimeObject *int64 `json:"recovery_time_object,omitempty"`
}

func (o ClusterDisasterRecovery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterDisasterRecovery struct{}"
	}

	return strings.Join([]string{"ClusterDisasterRecovery", string(data)}, " ")
}
