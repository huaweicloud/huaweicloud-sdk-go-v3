package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
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
