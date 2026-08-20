package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchBaselineIssueResponseResult 批量基线工作项的结果项
type BatchBaselineIssueResponseResult struct {

	// 基线成功的工作项列表。
	Success *[]IssueBaselineResult `json:"success,omitempty"`

	// 基线失败的工作项列表。
	Failed *[]IssueBaselineResult `json:"failed,omitempty"`

	// 成功数量。
	SuccessNum *int32 `json:"success_num,omitempty"`

	// 失败数量。
	FailNum *int32 `json:"fail_num,omitempty"`
}

func (o BatchBaselineIssueResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBaselineIssueResponseResult struct{}"
	}

	return strings.Join([]string{"BatchBaselineIssueResponseResult", string(data)}, " ")
}
