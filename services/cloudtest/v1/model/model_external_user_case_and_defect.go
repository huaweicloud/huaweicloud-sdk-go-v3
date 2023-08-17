package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExternalUserCaseAndDefect 用户创建用例以及关联缺陷信息
type ExternalUserCaseAndDefect struct {
	Creator *NameAndIdVo `json:"creator,omitempty"`

	// 缺陷数
	DefectCount *int32 `json:"defect_count,omitempty"`

	// 缺陷ID列表
	DefectIds *[]string `json:"defect_ids,omitempty"`

	// 用例ID
	TestcaseId *string `json:"testcase_id,omitempty"`

	// 分支ID
	BranchId *string `json:"branch_id,omitempty"`
}

func (o ExternalUserCaseAndDefect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalUserCaseAndDefect struct{}"
	}

	return strings.Join([]string{"ExternalUserCaseAndDefect", string(data)}, " ")
}
