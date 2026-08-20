package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateProcessInstanceReqCos struct {

	// 工作项编号
	Number *string `json:"number,omitempty"`

	// 工作项ID
	IssueId *string `json:"issue_id,omitempty"`

	// 工作项类型
	IssueCategory *string `json:"issue_category,omitempty"`

	// 变更类型
	ChangeType *string `json:"change_type,omitempty"`

	// 变更前
	BeforeChange *string `json:"before_change,omitempty"`

	// 变更后
	AfterChange *string `json:"after_change,omitempty"`
}

func (o CreateProcessInstanceReqCos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProcessInstanceReqCos struct{}"
	}

	return strings.Join([]string{"CreateProcessInstanceReqCos", string(data)}, " ")
}
