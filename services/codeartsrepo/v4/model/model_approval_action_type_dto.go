package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApprovalActionTypeDto struct {

	// **参数解释：** 审核/检视操作动作 - approve，审核通过 - reject，审核拒绝 - complete，检视通过 - reset，重置审核/检视结果
	ActionType *string `json:"action_type,omitempty"`

	// **参数解释：** 审核备注 **取值范围：** 不涉及。
	ApproverComment *string `json:"approver_comment,omitempty"`
}

func (o ApprovalActionTypeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApprovalActionTypeDto struct{}"
	}

	return strings.Join([]string{"ApprovalActionTypeDto", string(data)}, " ")
}
