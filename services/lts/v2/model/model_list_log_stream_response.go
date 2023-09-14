package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogStreamResponse Response Object
type ListLogStreamResponse struct {
	LogStreams     *[]LogStreamResBody `json:"log_streams,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListLogStreamResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogStreamResponse struct{}"
	}

	return strings.Join([]string{"ListLogStreamResponse", string(data)}, " ")
}
