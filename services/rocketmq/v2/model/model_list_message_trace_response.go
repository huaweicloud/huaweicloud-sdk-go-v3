package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMessageTraceResponse Response Object
type ListMessageTraceResponse struct {

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
