package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIssueCompletionRateRequest Request Object
type ShowIssueCompletionRateRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`
}

func (o ShowIssueCompletionRateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIssueCompletionRateRequest struct{}"
	}

	return strings.Join([]string{"ShowIssueCompletionRateRequest", string(data)}, " ")
}
