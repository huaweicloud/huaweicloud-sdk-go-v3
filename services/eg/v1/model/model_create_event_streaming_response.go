package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEventStreamingResponse Response Object
type CreateEventStreamingResponse struct {

	// 事件流ID
	EventStreamingID *string `json:"eventStreamingID,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEventStreamingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventStreamingResponse struct{}"
	}

	return strings.Join([]string{"CreateEventStreamingResponse", string(data)}, " ")
}
