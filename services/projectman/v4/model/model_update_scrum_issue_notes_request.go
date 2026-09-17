package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScrumIssueNotesRequest Request Object
type UpdateScrumIssueNotesRequest struct {
	Body *AddCommentsRequest `json:"body,omitempty"`
}

func (o UpdateScrumIssueNotesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScrumIssueNotesRequest struct{}"
	}

	return strings.Join([]string{"UpdateScrumIssueNotesRequest", string(data)}, " ")
}
