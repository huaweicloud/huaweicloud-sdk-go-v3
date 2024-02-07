package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalEipCountFilterTagsRequest Request Object
type ListGlobalEipCountFilterTagsRequest struct {
	Body *ListResourcesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListGlobalEipCountFilterTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalEipCountFilterTagsRequest struct{}"
	}

	return strings.Join([]string{"ListGlobalEipCountFilterTagsRequest", string(data)}, " ")
}
