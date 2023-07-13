package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEventStreamingRequest Request Object
type ShowEventStreamingRequest struct {

	// 事件流ID
	EventstreamingId string `json:"eventstreaming_id"`
}

func (o ShowEventStreamingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventStreamingRequest struct{}"
	}

	return strings.Join([]string{"ShowEventStreamingRequest", string(data)}, " ")
}
