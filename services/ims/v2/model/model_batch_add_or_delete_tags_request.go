package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchAddOrDeleteTagsRequest struct {
	// 镜像ID。

	ImageId string `json:"image_id"`

	Body *BatchAddOrDeleteTagsRequestBody `json:"body,omitempty"`
}

func (o BatchAddOrDeleteTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddOrDeleteTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchAddOrDeleteTagsRequest", string(data)}, " ")
}
