package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesByTagsRequest Request Object
type ListResourcesByTagsRequest struct {
	Body *ListResourceInstancesReq `json:"body,omitempty"`
}

func (o ListResourcesByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListResourcesByTagsRequest", string(data)}, " ")
}
