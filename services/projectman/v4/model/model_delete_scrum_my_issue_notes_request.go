package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScrumMyIssueNotesRequest Request Object
type DeleteScrumMyIssueNotesRequest struct {
	Body *DeleteIssueNoteParam `json:"body,omitempty"`
}

func (o DeleteScrumMyIssueNotesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScrumMyIssueNotesRequest struct{}"
	}

	return strings.Join([]string{"DeleteScrumMyIssueNotesRequest", string(data)}, " ")
}
