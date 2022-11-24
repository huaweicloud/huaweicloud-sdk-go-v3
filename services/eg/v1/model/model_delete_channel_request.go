package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteChannelRequest struct {

	// 指定查询的事件通道ID
	ChannelId string `json:"channel_id"`
}

func (o DeleteChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteChannelRequest struct{}"
	}

	return strings.Join([]string{"DeleteChannelRequest", string(data)}, " ")
}
