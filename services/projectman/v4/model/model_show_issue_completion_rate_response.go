package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowIssueCompletionRateResponse struct {

	// 不同类型的工作项完成率
	IssueCompletionRates *[]IssueCompletionRateV4IssueCompletionRates `json:"issue_completion_rates,omitempty" xml:"issue_completion_rates"`

	// 总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowIssueCompletionRateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIssueCompletionRateResponse struct{}"
	}

	return strings.Join([]string{"ShowIssueCompletionRateResponse", string(data)}, " ")
}
