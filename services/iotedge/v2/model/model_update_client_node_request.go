package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClientNodeRequest Request Object
type UpdateClientNodeRequest struct {

	// 边缘推送通道ID
	ChannelId string `json:"channel_id"`

	// 边缘节点ID
	NodeId string `json:"node_id"`

	Body *UpdateNodeChannelRequestDto `json:"body,omitempty"`
}

func (o UpdateClientNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClientNodeRequest struct{}"
	}

	return strings.Join([]string{"UpdateClientNodeRequest", string(data)}, " ")
}
