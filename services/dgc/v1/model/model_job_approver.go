package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobApprover struct {

	// 审批人名称
	ApproverName string `json:"approverName"`
}

func (o JobApprover) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobApprover struct{}"
	}

	return strings.Join([]string{"JobApprover", string(data)}, " ")
}
