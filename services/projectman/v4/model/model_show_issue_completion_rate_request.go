package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowIssueCompletionRateRequest struct {
	// devcloud的项目id

	ProjectId string `json:"project_id"`
}

func (o ShowIssueCompletionRateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIssueCompletionRateRequest struct{}"
	}

	return strings.Join([]string{"ShowIssueCompletionRateRequest", string(data)}, " ")
}
