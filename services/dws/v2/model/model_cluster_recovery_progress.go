package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterRecoveryProgress **参数解释**： 集群容灾进度详情。 **取值范围**： 不涉及。
type ClusterRecoveryProgress struct {

	// **参数解释**： 本次备份恢复ID。 **取值范围**： 不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释**： 动作类型。 **取值范围**： 不涉及。
	ActionType *string `json:"action_type,omitempty"`

	// **参数解释**： 待恢复的备份集ID。 **取值范围**： 不涉及。
	UnrestoreKeys *string `json:"unrestore_keys,omitempty"`

	// **参数解释**： 当前动作开始时间。 **取值范围**： 不涉及。
	ActionStartTime *string `json:"action_start_time,omitempty"`

	// **参数解释**： 当前动作结束时间。 **取值范围**： 不涉及。
	ActionEndTime *string `json:"action_end_time,omitempty"`
}

func (o ClusterRecoveryProgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterRecoveryProgress struct{}"
	}

	return strings.Join([]string{"ClusterRecoveryProgress", string(data)}, " ")
}
