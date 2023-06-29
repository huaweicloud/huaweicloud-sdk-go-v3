package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResumeEventStreamingResponse Response Object
type ResumeEventStreamingResponse struct {

	// 事件流ID
	EventStreamingID *string `json:"eventStreamingID,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResumeEventStreamingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumeEventStreamingResponse struct{}"
	}

	return strings.Join([]string{"ResumeEventStreamingResponse", string(data)}, " ")
}
