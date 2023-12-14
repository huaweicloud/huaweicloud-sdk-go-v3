package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStreamTagsRequest Request Object
type ShowStreamTagsRequest struct {

	// 通道ID。
	StreamId string `json:"stream_id"`
}

func (o ShowStreamTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStreamTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowStreamTagsRequest", string(data)}, " ")
}
