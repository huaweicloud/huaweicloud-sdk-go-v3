package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePushChannelRequest Request Object
type DeletePushChannelRequest struct {

	// 边缘推送通道ID
	ChannelId string `json:"channel_id"`
}

func (o DeletePushChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePushChannelRequest struct{}"
	}

	return strings.Join([]string{"DeletePushChannelRequest", string(data)}, " ")
}
