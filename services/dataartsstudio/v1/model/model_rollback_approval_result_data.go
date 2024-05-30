package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollbackApprovalResultData 审批单撤回的返回结果，成功的个数。
type RollbackApprovalResultData struct {

	// 审批单撤回的对象个数
	Value *string `json:"value,omitempty"`
}

func (o RollbackApprovalResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackApprovalResultData struct{}"
	}

	return strings.Join([]string{"RollbackApprovalResultData", string(data)}, " ")
}
