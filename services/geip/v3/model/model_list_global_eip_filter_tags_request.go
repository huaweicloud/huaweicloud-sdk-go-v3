package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalEipFilterTagsRequest Request Object
type ListGlobalEipFilterTagsRequest struct {
	Limit *[]int32 `json:"limit,omitempty"`

	Offset *[]int32 `json:"offset,omitempty"`

	Body *ListResourcesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListGlobalEipFilterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalEipFilterTagsRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalEipFilterTagsRequest", string(data)}, " ")
}
