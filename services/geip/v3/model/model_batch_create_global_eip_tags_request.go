package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateGlobalEipTagsRequest Request Object
type BatchCreateGlobalEipTagsRequest struct {
	ResourceId string `json:"resource_id"`

	Body *BatchCreateV2RequestBody `json:"body,omitempty"`
}

func (o BatchCreateGlobalEipTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateGlobalEipTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateGlobalEipTagsRequest", string(data)}, " ")
}
