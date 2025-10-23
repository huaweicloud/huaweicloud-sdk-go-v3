package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCommitDiscussionsResponse Response Object
type ListCommitDiscussionsResponse struct {
	Body           *[]CommitDiscussionDto `json:"body,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListCommitDiscussionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommitDiscussionsResponse struct{}"
	}

	return strings.Join([]string{"ListCommitDiscussionsResponse", string(data)}, " ")
}
