package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Approver 审批人
type Approver struct {

	// 审批人名称
	ApproverName string `json:"approver_name"`
}

func (o Approver) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Approver struct{}"
	}

	return strings.Join([]string{"Approver", string(data)}, " ")
}
