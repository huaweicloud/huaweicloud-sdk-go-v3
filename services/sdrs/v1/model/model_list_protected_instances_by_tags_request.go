package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListProtectedInstancesByTagsRequest struct {
	Body *ListProtectedInstancesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListProtectedInstancesByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectedInstancesByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProtectedInstancesByTagsRequest", string(data)}, " ")
}
