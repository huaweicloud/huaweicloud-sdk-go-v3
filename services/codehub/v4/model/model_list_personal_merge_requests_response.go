package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPersonalMergeRequestsResponse Response Object
type ListPersonalMergeRequestsResponse struct {
	Body           *[]MergeRequestListBasicDto `json:"body,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListPersonalMergeRequestsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPersonalMergeRequestsResponse struct{}"
	}

	return strings.Join([]string{"ListPersonalMergeRequestsResponse", string(data)}, " ")
}
