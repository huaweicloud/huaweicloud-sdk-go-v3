package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBranchRelatedWorkItemsResponse Response Object
type ListBranchRelatedWorkItemsResponse struct {
	Body           *[]WorkItemSimpleDto `json:"body,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListBranchRelatedWorkItemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBranchRelatedWorkItemsResponse struct{}"
	}

	return strings.Join([]string{"ListBranchRelatedWorkItemsResponse", string(data)}, " ")
}
