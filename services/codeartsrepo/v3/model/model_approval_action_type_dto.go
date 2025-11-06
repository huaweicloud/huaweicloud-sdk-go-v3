package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApprovalActionTypeDto struct {

	// 操作类型,取值,通过:approve; 拒绝:reject; 撤销:reset
	ActionType *string `json:"action_type,omitempty"`
}

func (o ApprovalActionTypeDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApprovalActionTypeDto struct{}"
	}

	return strings.Join([]string{"ApprovalActionTypeDto", string(data)}, " ")
}
