package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScrumMyIssueNotesRequest Request Object
type UpdateScrumMyIssueNotesRequest struct {
	Body *UpdateCommentsRequest `json:"body,omitempty"`
}

func (o UpdateScrumMyIssueNotesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScrumMyIssueNotesRequest struct{}"
	}

	return strings.Join([]string{"UpdateScrumMyIssueNotesRequest", string(data)}, " ")
}
