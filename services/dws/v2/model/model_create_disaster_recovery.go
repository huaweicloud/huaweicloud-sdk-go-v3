package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDisasterRecovery struct {

	// 名称
	Name string `json:"name"`

	// 容灾类型
	DrType string `json:"dr_type"`

	// 主集群ID
	PrimaryClusterId string `json:"primary_cluster_id"`

	// 备集群ID
	StandbyClusterId string `json:"standby_cluster_id"`

	// 同步周期
	DrSyncPeriod string `json:"dr_sync_period"`

	// 主集群OBS桶
	PrimaryObsBucket *string `json:"primary_obs_bucket,omitempty"`

	// 备集群obs桶
	StandbyObsBucket *string `json:"standby_obs_bucket,omitempty"`
}

func (o CreateDisasterRecovery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDisasterRecovery struct{}"
	}

	return strings.Join([]string{"CreateDisasterRecovery", string(data)}, " ")
}
