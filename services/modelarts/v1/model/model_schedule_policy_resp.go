package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SchedulePolicyResp 训练作业调度策略
type SchedulePolicyResp struct {
	RequiredAffinity *RequiredAffinityResp `json:"required_affinity,omitempty"`

	// **参数解释**：训练作业优先级。 **取值范围**：0-3
	Priority *int32 `json:"priority,omitempty"`

	// **参数解释**：是否可以被抢占。 **取值范围**： - true：可以被抢占 - false：不可以被抢占
	Preemptible *bool `json:"preemptible,omitempty"`
}

func (o SchedulePolicyResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchedulePolicyResp struct{}"
	}

	return strings.Join([]string{"SchedulePolicyResp", string(data)}, " ")
}
