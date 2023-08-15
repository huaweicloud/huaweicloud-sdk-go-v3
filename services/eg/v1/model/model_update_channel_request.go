package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateChannelRequest Request Object
type UpdateChannelRequest struct {

	// 指定查询的事件通道ID
	ChannelId string `json:"channel_id"`

	Body *ChannelUpdateReq `json:"body,omitempty"`
}

func (o UpdateChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateChannelRequest struct{}"
	}

	return strings.Join([]string{"UpdateChannelRequest", string(data)}, " ")
}
