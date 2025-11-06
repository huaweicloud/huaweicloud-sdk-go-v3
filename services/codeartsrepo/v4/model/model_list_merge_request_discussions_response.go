package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestDiscussionsResponse Response Object
type ListMergeRequestDiscussionsResponse struct {
	Body           *[]MergeRequestDiscussionDto `json:"body,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListMergeRequestDiscussionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestDiscussionsResponse struct{}"
	}

	return strings.Join([]string{"ListMergeRequestDiscussionsResponse", string(data)}, " ")
}
