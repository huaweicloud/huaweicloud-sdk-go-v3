package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailOfChannelRequest struct {

	// 指定查询的事件通道ID
	ChannelId string `json:"channel_id" xml:"channel_id"`
}

func (o ShowDetailOfChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfChannelRequest struct{}"
	}

	return strings.Join([]string{"ShowDetailOfChannelRequest", string(data)}, " ")
}
