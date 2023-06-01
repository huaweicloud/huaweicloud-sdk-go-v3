package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteEventStreamingResponse struct {

	// 事件流ID
	EventStreamingID *string `json:"eventStreamingID,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteEventStreamingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEventStreamingResponse struct{}"
	}

	return strings.Join([]string{"DeleteEventStreamingResponse", string(data)}, " ")
}
