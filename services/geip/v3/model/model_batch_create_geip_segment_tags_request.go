package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateGeipSegmentTagsRequest Request Object
type BatchCreateGeipSegmentTagsRequest struct {
	ResourceId string `json:"resource_id"`

	Body *BatchCreateV2RequestBody `json:"body,omitempty"`
}

func (o BatchCreateGeipSegmentTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateGeipSegmentTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateGeipSegmentTagsRequest", string(data)}, " ")
}
