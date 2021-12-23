package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListInstancesByTagsRequest struct {
	Body *ListInstancesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListInstancesByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesByTagsRequest", string(data)}, " ")
}
