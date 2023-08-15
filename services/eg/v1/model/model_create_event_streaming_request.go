package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEventStreamingRequest Request Object
type CreateEventStreamingRequest struct {
	Body *EventStreamingCreateReq `json:"body,omitempty"`
}

func (o CreateEventStreamingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventStreamingRequest struct{}"
	}

	return strings.Join([]string{"CreateEventStreamingRequest", string(data)}, " ")
}
