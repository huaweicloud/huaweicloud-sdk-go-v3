package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CocUpdateChangeRequestBodyTicketInfo 变更主单信息。
type CocUpdateChangeRequestBodyTicketInfo struct {

	// -| 工单操作的类型，枚举值。 · phase_change_end：完成 · phase_change_cancel：撤销 · phase_change_draft：草稿 · phase_change_implemente：变更实施和验证 · phase_change_apply：申请人确认 · phase_change_approve：审批中 · phase_change_close：关闭
	Phase *string `json:"phase,omitempty"`

	// 工单状态。
	WorkFlowStatus *string `json:"work_flow_status,omitempty"`
}

func (o CocUpdateChangeRequestBodyTicketInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocUpdateChangeRequestBodyTicketInfo struct{}"
	}

	return strings.Join([]string{"CocUpdateChangeRequestBodyTicketInfo", string(data)}, " ")
}
