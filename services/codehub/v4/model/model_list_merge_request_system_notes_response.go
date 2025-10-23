package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestSystemNotesResponse Response Object
type ListMergeRequestSystemNotesResponse struct {
	Body           *[]NoteDto `json:"body,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListMergeRequestSystemNotesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestSystemNotesResponse struct{}"
	}

	return strings.Join([]string{"ListMergeRequestSystemNotesResponse", string(data)}, " ")
}
