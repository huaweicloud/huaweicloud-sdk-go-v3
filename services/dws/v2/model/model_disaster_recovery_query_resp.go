package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisasterRecoveryQueryResp **参数解释**： 查询容灾信息返回体。 **取值范围**： 不涉及。
type DisasterRecoveryQueryResp struct {

	// **参数解释**： 容灾ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 容灾名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 容灾类型。 **取值范围**： 不涉及。
	DrType *string `json:"dr_type,omitempty"`

	// **参数解释**： 容灾状态。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	PrimaryCluster *DisasterRecoveryCluster `json:"primary_cluster,omitempty"`

	StandbyCluster *DisasterRecoveryCluster `json:"standby_cluster,omitempty"`

	// **参数解释**： 容灾同步周期。 **取值范围**： 不涉及。
	DrSyncPeriod *string `json:"dr_sync_period,omitempty"`

	// **参数解释**： 容灾启动时间。 **取值范围**： 不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**： 容灾创建时间。 **取值范围**： 不涉及。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o DisasterRecoveryQueryResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisasterRecoveryQueryResp struct{}"
	}

	return strings.Join([]string{"DisasterRecoveryQueryResp", string(data)}, " ")
}
