package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueCompletionRateV4IssueCompletionRates struct {
	IssueStatus *IssueCompletionRateV4IssueStatus `json:"issue_status,omitempty" xml:"issue_status"`

	// 工作项类型id,1需求,2任务/Task,3缺陷/Bug,5Epic,6Feature,7Story
	TrackerId *int32 `json:"tracker_id,omitempty" xml:"tracker_id"`
}

func (o IssueCompletionRateV4IssueCompletionRates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueCompletionRateV4IssueCompletionRates struct{}"
	}

	return strings.Join([]string{"IssueCompletionRateV4IssueCompletionRates", string(data)}, " ")
}
