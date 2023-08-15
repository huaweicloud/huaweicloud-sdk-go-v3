package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutEventsRequest Request Object
type PutEventsRequest struct {

	// 指定查询的事件通道ID
	ChannelId string `json:"channel_id"`

	Body *PutEventsReq `json:"body,omitempty"`
}

func (o PutEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutEventsRequest struct{}"
	}

	return strings.Join([]string{"PutEventsRequest", string(data)}, " ")
}
