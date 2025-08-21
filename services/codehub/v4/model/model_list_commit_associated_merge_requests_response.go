package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCommitAssociatedMergeRequestsResponse Response Object
type ListCommitAssociatedMergeRequestsResponse struct {
	Body           *[]CommitMergeRequestDto `json:"body,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListCommitAssociatedMergeRequestsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCommitAssociatedMergeRequestsResponse struct{}"
	}

	return strings.Join([]string{"ListCommitAssociatedMergeRequestsResponse", string(data)}, " ")
}
