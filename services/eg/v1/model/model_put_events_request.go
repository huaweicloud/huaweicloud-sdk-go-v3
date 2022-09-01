package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type PutEventsRequest struct {

	// 指定查询的事件通道ID
	ChannelId string `json:"channel_id" xml:"channel_id"`

	Body *PutEventsReq `json:"body,omitempty" xml:"body"`
}

func (o PutEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutEventsRequest struct{}"
	}

	return strings.Join([]string{"PutEventsRequest", string(data)}, " ")
}
