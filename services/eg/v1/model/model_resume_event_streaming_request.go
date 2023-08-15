package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResumeEventStreamingRequest Request Object
type ResumeEventStreamingRequest struct {

	// 事件流ID
	EventstreamingId string `json:"eventstreaming_id"`

	Body *EventStreamingOperateReq `json:"body,omitempty"`
}

func (o ResumeEventStreamingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumeEventStreamingRequest struct{}"
	}

	return strings.Join([]string{"ResumeEventStreamingRequest", string(data)}, " ")
}
