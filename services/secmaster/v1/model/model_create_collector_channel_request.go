package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCollectorChannelRequest Request Object
type CreateCollectorChannelRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *CreateChannelDto `json:"body,omitempty"`
}

func (o CreateCollectorChannelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCollectorChannelRequest struct{}"
	}

	return strings.Join([]string{"CreateCollectorChannelRequest", string(data)}, " ")
}
