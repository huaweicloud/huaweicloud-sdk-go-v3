package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCollectorChannelRequest Request Object
type ShowCollectorChannelRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 采集通道ID
	ChannelId string `json:"channel_id"`
}

func (o ShowCollectorChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCollectorChannelRequest struct{}"
	}

	return strings.Join([]string{"ShowCollectorChannelRequest", string(data)}, " ")
}
