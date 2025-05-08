package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMessageTraceResponse Response Object
type ListMessageTraceResponse struct {

	// 总数。
	Total float32 `json:"total,omitempty"`

	// 下个分页的offset。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// 上个分页的offset。
	PreviousOffset *int32 `json:"previous_offset,omitempty"`

	// 消息轨迹列表。
	Trace          *[]ListMessageTraceRespTrace `json:"trace,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListMessageTraceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageTraceResponse struct{}"
	}

	return strings.Join([]string{"ListMessageTraceResponse", string(data)}, " ")
}
