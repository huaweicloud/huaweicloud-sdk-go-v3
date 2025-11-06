package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMergeRequestCommentsByLineResponse Response Object
type ShowMergeRequestCommentsByLineResponse struct {
	Body           *[]CommentPathDto `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowMergeRequestCommentsByLineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMergeRequestCommentsByLineResponse struct{}"
	}

	return strings.Join([]string{"ShowMergeRequestCommentsByLineResponse", string(data)}, " ")
}
