package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProtectedInstancesByTagsRequest Request Object
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
