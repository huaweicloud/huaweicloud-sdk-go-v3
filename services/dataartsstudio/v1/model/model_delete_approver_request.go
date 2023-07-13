package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApproverRequest Request Object
type DeleteApproverRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 审批人id
	ApproverIds string `json:"approver_ids"`
}

func (o DeleteApproverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApproverRequest struct{}"
	}

	return strings.Join([]string{"DeleteApproverRequest", string(data)}, " ")
}
