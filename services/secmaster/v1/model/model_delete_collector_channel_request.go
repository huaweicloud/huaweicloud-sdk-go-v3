package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCollectorChannelRequest Request Object
type DeleteCollectorChannelRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 采集通道ID
	ChannelId string `json:"channel_id"`
}

func (o DeleteCollectorChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCollectorChannelRequest struct{}"
	}

	return strings.Join([]string{"DeleteCollectorChannelRequest", string(data)}, " ")
}
