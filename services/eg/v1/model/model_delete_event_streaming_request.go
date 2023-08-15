package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEventStreamingRequest Request Object
type DeleteEventStreamingRequest struct {

	// 事件流ID
	EventstreamingId string `json:"eventstreaming_id"`
}

func (o DeleteEventStreamingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEventStreamingRequest struct{}"
	}

	return strings.Join([]string{"DeleteEventStreamingRequest", string(data)}, " ")
}
