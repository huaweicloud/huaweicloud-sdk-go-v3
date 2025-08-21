package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryWorkItemsResponse Response Object
type ListRepositoryWorkItemsResponse struct {
	Body *[]ReqWorkItemDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRepositoryWorkItemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryWorkItemsResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryWorkItemsResponse", string(data)}, " ")
}
