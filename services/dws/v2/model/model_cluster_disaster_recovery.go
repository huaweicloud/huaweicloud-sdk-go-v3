package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterDisasterRecovery **参数解释**： 集群容灾恢复详情。 **取值范围**： 不涉及。
type ClusterDisasterRecovery struct {
	PrimaryCluster *ClusterRecoveryProgress `json:"primary_cluster,omitempty"`

	StandbyCluster *ClusterRecoveryProgress `json:"standby_cluster,omitempty"`

	// **参数解释**： 故障发生时间。 **取值范围**： 不涉及。
	LatestBarrierTime *string `json:"latest_barrier_time,omitempty"`

	// **参数解释**： 上一个备份集恢复消耗时间，单位：秒（s）。 **取值范围**： 不涉及。
	LastRecoverySpend *int64 `json:"last_recovery_spend,omitempty"`

	// **参数解释**： 数据恢复目标时间，单位：秒（s）。 **取值范围**： 不涉及。
	RecoveryPointObject *int64 `json:"recovery_point_object,omitempty"`

	// **参数解释**： 服务恢复目标时间，单位：秒（s）。 **取值范围**： 不涉及。
	RecoveryTimeObject *int64 `json:"recovery_time_object,omitempty"`
}

func (o ClusterDisasterRecovery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterDisasterRecovery struct{}"
	}

	return strings.Join([]string{"ClusterDisasterRecovery", string(data)}, " ")
}
