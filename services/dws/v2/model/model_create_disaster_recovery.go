package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDisasterRecovery struct {

	// **参数解释**： 名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 容灾类型。 **取值范围**： 不涉及。
	DrType string `json:"dr_type"`

	// **参数解释**： 主集群ID。 **取值范围**： 不涉及。
	PrimaryClusterId string `json:"primary_cluster_id"`

	// **参数解释**： 备集群ID。 **取值范围**： 不涉及。
	StandbyClusterId string `json:"standby_cluster_id"`

	// **参数解释**： 同步周期。 **取值范围**： 不涉及。
	DrSyncPeriod string `json:"dr_sync_period"`

	// **参数解释**： 主集群OBS桶。 **取值范围**： 不涉及。
	PrimaryObsBucket *string `json:"primary_obs_bucket,omitempty"`

	// **参数解释**： 备集群obs桶。 **取值范围**： 不涉及。
	StandbyObsBucket *string `json:"standby_obs_bucket,omitempty"`
}

func (o CreateDisasterRecovery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDisasterRecovery struct{}"
	}

	return strings.Join([]string{"CreateDisasterRecovery", string(data)}, " ")
}
