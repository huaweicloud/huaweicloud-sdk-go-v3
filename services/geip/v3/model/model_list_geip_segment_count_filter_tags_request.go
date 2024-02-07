package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGeipSegmentCountFilterTagsRequest Request Object
type ListGeipSegmentCountFilterTagsRequest struct {
	Body *ListResourcesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListGeipSegmentCountFilterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeipSegmentCountFilterTagsRequest struct{}"
	}

	return strings.Join([]string{"ListGeipSegmentCountFilterTagsRequest", string(data)}, " ")
}
