package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClientNodeRequest Request Object
type CreateClientNodeRequest struct {

	// 边缘推送通道ID
	ChannelId string `json:"channel_id"`

	Body *CreateClientNodeRequestDto `json:"body,omitempty"`
}

func (o CreateClientNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClientNodeRequest struct{}"
	}

	return strings.Join([]string{"CreateClientNodeRequest", string(data)}, " ")
}
