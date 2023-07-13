package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateNodeTagsRequest Request Object
type BatchCreateNodeTagsRequest struct {
	Body *MultiResourcesMultiTags `json:"body,omitempty"`
}

func (o BatchCreateNodeTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateNodeTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateNodeTagsRequest", string(data)}, " ")
}
