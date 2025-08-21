package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestCommitsResponse Response Object
type ListMergeRequestCommitsResponse struct {
	Body           *[]MergeRequestCommitDto `json:"body,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListMergeRequestCommitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestCommitsResponse struct{}"
	}

	return strings.Join([]string{"ListMergeRequestCommitsResponse", string(data)}, " ")
}
