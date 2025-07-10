package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CocTicketDetailInfoResponseDataSubTickets struct {

	// 变更单单号。
	ChangeTicketId *string `json:"change_ticket_id,omitempty"`

	// 变更子单单号。
	ChangeTicketIdSub *string `json:"change_ticket_id_sub,omitempty"`

	// 是否需要变更。
	WhetherToChange *bool `json:"whether_to_change,omitempty"`

	// 是否已删除。
	IsDeleted *bool `json:"is_deleted,omitempty"`

	// 变更工单ID
	Id *string `json:"id,omitempty"`

	// 变更工单主键ID。
	MainTicketId *string `json:"main_ticket_id,omitempty"`

	// 父级工单ID。
	ParentTicketId *string `json:"parent_ticket_id,omitempty"`

	// 问题单ID。
	TicketId *string `json:"ticket_id,omitempty"`

	// 问题单单号。
	RealTicketId *string `json:"real_ticket_id,omitempty"`

	// 工单路径。
	TicketPath *string `json:"ticket_path,omitempty"`

	// region信息。
	TargetValue *string `json:"target_value,omitempty"`

	// 子单类型。
	TargetType *string `json:"target_type,omitempty"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 创单人。
	Creator *string `json:"creator,omitempty"`

	// 操作人。
	Operator *string `json:"operator,omitempty"`
}

func (o CocTicketDetailInfoResponseDataSubTickets) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocTicketDetailInfoResponseDataSubTickets struct{}"
	}

	return strings.Join([]string{"CocTicketDetailInfoResponseDataSubTickets", string(data)}, " ")
}
