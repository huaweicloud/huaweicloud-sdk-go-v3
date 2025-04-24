package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEventStreamingResponse Response Object
type DeleteEventStreamingResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteEventStreamingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEventStreamingResponse struct{}"
	}

	return strings.Join([]string{"DeleteEventStreamingResponse", string(data)}, " ")
}
