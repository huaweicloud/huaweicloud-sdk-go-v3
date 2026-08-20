package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueBaselineResult 工作项基线返回值
type IssueBaselineResult struct {

	// 变更的工作项ID。
	Id *string `json:"id,omitempty"`

	// 工作项变更人ID。
	ModifiedBy *string `json:"modified_by,omitempty"`

	// 工作项基线结果。
	Baseline *string `json:"baseline,omitempty"`

	// 工作项基线的操作记录ID。
	OperationId *string `json:"operation_id,omitempty"`

	// 工作项完成基线的unix时间戳，单位：毫秒。
	ModifiedDate *string `json:"modified_date,omitempty"`

	// 基线的工作项编号。 基线失败时返回。
	Number *string `json:"number,omitempty"`

	// 基线的工作项标题。 基线失败时返回。
	Title *string `json:"title,omitempty"`

	// 工作项基线失败原因。 基线失败时返回。
	FailMessage *string `json:"fail_message,omitempty"`
}

func (o IssueBaselineResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueBaselineResult struct{}"
	}

	return strings.Join([]string{"IssueBaselineResult", string(data)}, " ")
}
