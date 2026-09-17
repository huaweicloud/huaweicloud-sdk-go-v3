package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScrumMyIssueNotesResponse Response Object
type UpdateScrumMyIssueNotesResponse struct {
	Result *UpdateNoteResponseResult `json:"result,omitempty"`

	// **参数解释：** 返回状态。 **取值范围：** success：返回成功。 error：返回失败。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateScrumMyIssueNotesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScrumMyIssueNotesResponse struct{}"
	}

	return strings.Join([]string{"UpdateScrumMyIssueNotesResponse", string(data)}, " ")
}
