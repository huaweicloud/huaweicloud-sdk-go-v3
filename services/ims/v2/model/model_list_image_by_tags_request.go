package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageByTagsRequest Request Object
type ListImageByTagsRequest struct {
	Body *ListImageByTagsRequestBody `json:"body,omitempty"`
}

func (o ListImageByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListImageByTagsRequest", string(data)}, " ")
}
