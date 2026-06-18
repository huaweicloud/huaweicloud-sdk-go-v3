package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectMergeRequestsResponse Response Object
type ListProjectMergeRequestsResponse struct {
	Body *[]MergeRequestListBasicDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListProjectMergeRequestsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectMergeRequestsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectMergeRequestsResponse", string(data)}, " ")
}
