package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCommitCommentsByLineResponse Response Object
type ShowCommitCommentsByLineResponse struct {
	Body           *[]CommentPathDto `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowCommitCommentsByLineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommitCommentsByLineResponse struct{}"
	}

	return strings.Join([]string{"ShowCommitCommentsByLineResponse", string(data)}, " ")
}
