package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScrumMyIssueNotesResponse Response Object
type DeleteScrumMyIssueNotesResponse struct {
	Result *DeleteIssueNoteResultResult `json:"result,omitempty"`

	// **参数解释**： 接口整体响应状态。 **取值范围**： - success：接口请求成功。 - error：接口请求失败。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteScrumMyIssueNotesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScrumMyIssueNotesResponse struct{}"
	}

	return strings.Join([]string{"DeleteScrumMyIssueNotesResponse", string(data)}, " ")
}
