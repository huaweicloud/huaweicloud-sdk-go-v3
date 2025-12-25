package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCollectorChannelOperationRequest Request Object
type CreateCollectorChannelOperationRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 采集通道ID
	ChannelId string `json:"channel_id"`

	Body *OperationDto `json:"body,omitempty"`
}

func (o CreateCollectorChannelOperationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCollectorChannelOperationRequest struct{}"
	}

	return strings.Join([]string{"CreateCollectorChannelOperationRequest", string(data)}, " ")
}
