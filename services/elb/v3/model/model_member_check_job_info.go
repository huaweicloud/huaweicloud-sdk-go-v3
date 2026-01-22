package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MemberCheckJobInfo **参数解释**：后端服务器检查任务结果。
type MemberCheckJobInfo struct {

	// **参数解释**：检查任务状态。  **取值范围**： - processed：检查任务执行完成。 - processing：检查任务执行中。 - failed：检查任务执行失败。
	Status *string `json:"status,omitempty"`

	Result *MemberCheckJobResult `json:"result,omitempty"`

	// **参数解释**：任务创建时间。  **取值范围**：不涉及
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**：任务更新时间。  **取值范围**：不涉及
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释**：任务ID。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**：检查项总个数。  **取值范围**：不涉及
	CheckItemTotalNum *int32 `json:"check_item_total_num,omitempty"`

	// **参数解释**：已检查完成的检查项个数。  **取值范围**：不涉及
	CheckItemFinishedNum *int32 `json:"check_item_finished_num,omitempty"`

	// **参数解释**：后端服务器所关联的监听器，查询在该监听器下后端服务器的状态。  **取值范围**：不涉及
	ListenerId *string `json:"listener_id,omitempty"`

	// **参数解释**：后端服务器ID。  **取值范围**：不涉及
	MemberId *string `json:"member_id,omitempty"`
}

func (o MemberCheckJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberCheckJobInfo struct{}"
	}

	return strings.Join([]string{"MemberCheckJobInfo", string(data)}, " ")
}
