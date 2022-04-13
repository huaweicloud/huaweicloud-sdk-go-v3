package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueCompletionRateV4IssueCompletionRates struct {
	IssueStatus *IssueCompletionRateV4IssueStatus `json:"issue_status,omitempty"`
	// 工作项类型id,1需求,2任务/task,3缺陷/bug,5epic,6feature,7story

	TrackerId *int32 `json:"tracker_id,omitempty"`
}

func (o IssueCompletionRateV4IssueCompletionRates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueCompletionRateV4IssueCompletionRates struct{}"
	}

	return strings.Join([]string{"IssueCompletionRateV4IssueCompletionRates", string(data)}, " ")
}
