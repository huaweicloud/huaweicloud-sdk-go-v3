package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogStreamsResponse Response Object
type ListLogStreamsResponse struct {

	// 日志流数组
	LogStreams     *[]ListLogStreamsResponseBody1LogStreams `json:"log_streams,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o ListLogStreamsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogStreamsResponse struct{}"
	}

	return strings.Join([]string{"ListLogStreamsResponse", string(data)}, " ")
}
