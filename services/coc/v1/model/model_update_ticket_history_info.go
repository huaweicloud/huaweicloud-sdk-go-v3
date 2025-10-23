package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTicketHistoryInfo struct {

	// 操作标识
	ActionId *string `json:"action_id,omitempty"`

	// 动作
	Action *string `json:"action,omitempty"`

	// 子动作
	SubAction *string `json:"sub_action,omitempty"`

	// 操作人
	Operator *string `json:"operator,omitempty"`

	// 评论
	Comment *string `json:"comment,omitempty"`
}

func (o UpdateTicketHistoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTicketHistoryInfo struct{}"
	}

	return strings.Join([]string{"UpdateTicketHistoryInfo", string(data)}, " ")
}
