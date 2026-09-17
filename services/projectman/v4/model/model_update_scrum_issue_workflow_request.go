package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScrumIssueWorkflowRequest Request Object
type UpdateScrumIssueWorkflowRequest struct {
	Body *IssueFlowRequest `json:"body,omitempty"`
}

func (o UpdateScrumIssueWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScrumIssueWorkflowRequest struct{}"
	}

	return strings.Join([]string{"UpdateScrumIssueWorkflowRequest", string(data)}, " ")
}
