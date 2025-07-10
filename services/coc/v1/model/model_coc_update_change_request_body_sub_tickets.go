package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CocUpdateChangeRequestBodySubTickets struct {

	// 子单ID。
	TicketId *string `json:"ticket_id,omitempty"`

	// 变更结果。
	ChangeResult *string `json:"change_result,omitempty"`

	// 在时间窗内是否可验证。
	IsVerifiedInChangeTime *bool `json:"is_verified_in_change_time,omitempty"`

	// 验证文档ID。
	VerifiedDocs *string `json:"verified_docs,omitempty"`

	// 变更失败原因描述。
	Comment *string `json:"comment,omitempty"`

	// 变更失败类型。
	ChangeFailType *string `json:"change_fail_type,omitempty"`

	// 回退开始时间。
	RollbackStartTime *int64 `json:"rollback_start_time,omitempty"`

	// 回退结束时间。
	RollbackEndTime *int64 `json:"rollback_end_time,omitempty"`

	// 是否回退成功。
	IsRollbackSuccess *bool `json:"is_rollback_success,omitempty"`

	// 是否被监控发现。
	IsMonitorFound *bool `json:"is_monitor_found,omitempty"`
}

func (o CocUpdateChangeRequestBodySubTickets) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocUpdateChangeRequestBodySubTickets struct{}"
	}

	return strings.Join([]string{"CocUpdateChangeRequestBodySubTickets", string(data)}, " ")
}
