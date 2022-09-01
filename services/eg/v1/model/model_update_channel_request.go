package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateChannelRequest struct {

	// 指定查询的事件通道ID
	ChannelId string `json:"channel_id" xml:"channel_id"`

	Body *ChannelUpdateReq `json:"body,omitempty" xml:"body"`
}

func (o UpdateChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateChannelRequest struct{}"
	}

	return strings.Join([]string{"UpdateChannelRequest", string(data)}, " ")
}
