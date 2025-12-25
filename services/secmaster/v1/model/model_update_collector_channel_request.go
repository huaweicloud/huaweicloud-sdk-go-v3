package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCollectorChannelRequest Request Object
type UpdateCollectorChannelRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 采集通道ID
	ChannelId string `json:"channel_id"`

	Body *CreateChannelDto `json:"body,omitempty"`
}

func (o UpdateCollectorChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCollectorChannelRequest struct{}"
	}

	return strings.Join([]string{"UpdateCollectorChannelRequest", string(data)}, " ")
}
