package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteGeipSegmentTagsRequest Request Object
type BatchDeleteGeipSegmentTagsRequest struct {
	ResourceId string `json:"resource_id"`

	Body *BatchDeleteV2RequestBody `json:"body,omitempty"`
}

func (o BatchDeleteGeipSegmentTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteGeipSegmentTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteGeipSegmentTagsRequest", string(data)}, " ")
}
