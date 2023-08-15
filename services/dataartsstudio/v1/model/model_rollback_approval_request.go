package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollbackApprovalRequest Request Object
type RollbackApprovalRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 审批单id
	Ids string `json:"ids"`
}

func (o RollbackApprovalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackApprovalRequest struct{}"
	}

	return strings.Join([]string{"RollbackApprovalRequest", string(data)}, " ")
}
