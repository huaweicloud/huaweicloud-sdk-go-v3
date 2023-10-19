package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudConnectionsByTagsRequest Request Object
type ListCloudConnectionsByTagsRequest struct {
	Body *ListCloudConnectionsByTagsRequestBody `json:"body,omitempty"`
}

func (o ListCloudConnectionsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudConnectionsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListCloudConnectionsByTagsRequest", string(data)}, " ")
}
