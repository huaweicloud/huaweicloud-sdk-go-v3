package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteReq struct {

	// 在开启审批开关后，需要填写该字段，表示作业（或脚本）审批人。
	Approvers *[]JobApprover `json:"approvers,omitempty"`
}

func (o DeleteReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReq struct{}"
	}

	return strings.Join([]string{"DeleteReq", string(data)}, " ")
}
