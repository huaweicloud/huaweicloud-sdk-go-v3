package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScrumIssueNotesResponse Response Object
type UpdateScrumIssueNotesResponse struct {
	Result *IssueInfoResponseResult `json:"result,omitempty"`

	// **参数解释：** 返回状态。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateScrumIssueNotesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScrumIssueNotesResponse struct{}"
	}

	return strings.Join([]string{"UpdateScrumIssueNotesResponse", string(data)}, " ")
}
