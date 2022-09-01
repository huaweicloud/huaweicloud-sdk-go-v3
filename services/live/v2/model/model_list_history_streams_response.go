package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHistoryStreamsResponse struct {

	// 历史流信息列表。
	HistoryStreamList *[]HistoryStreamInfo `json:"history_stream_list,omitempty" xml:"history_stream_list"`

	// 总记录数。
	Total *int32 `json:"total,omitempty" xml:"total"`

	XRequestId     *string `json:"X-Request-Id,omitempty" xml:"X-Request-Id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListHistoryStreamsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryStreamsResponse struct{}"
	}

	return strings.Join([]string{"ListHistoryStreamsResponse", string(data)}, " ")
}
