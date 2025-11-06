package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMergeRequestDiscussionNoteResponse Response Object
type CreateMergeRequestDiscussionNoteResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *NoteDto `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMergeRequestDiscussionNoteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMergeRequestDiscussionNoteResponse struct{}"
	}

	return strings.Join([]string{"CreateMergeRequestDiscussionNoteResponse", string(data)}, " ")
}
