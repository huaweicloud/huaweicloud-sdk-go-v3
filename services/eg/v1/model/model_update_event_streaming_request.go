package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateEventStreamingRequest struct {

	// 事件流ID
	EventstreamingId string `json:"eventstreaming_id"`

	Body *EventStreamingUpdateReq `json:"body,omitempty"`
}

func (o UpdateEventStreamingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEventStreamingRequest struct{}"
	}

	return strings.Join([]string{"UpdateEventStreamingRequest", string(data)}, " ")
}
