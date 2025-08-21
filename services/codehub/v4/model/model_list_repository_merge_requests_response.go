package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryMergeRequestsResponse Response Object
type ListRepositoryMergeRequestsResponse struct {
	Body           *[]MergeRequestListBasicDto `json:"body,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListRepositoryMergeRequestsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryMergeRequestsResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryMergeRequestsResponse", string(data)}, " ")
}
