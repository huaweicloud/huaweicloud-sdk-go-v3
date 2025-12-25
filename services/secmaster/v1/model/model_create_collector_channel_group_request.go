package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCollectorChannelGroupRequest Request Object
type CreateCollectorChannelGroupRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *GroupDto `json:"body,omitempty"`
}

func (o CreateCollectorChannelGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCollectorChannelGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateCollectorChannelGroupRequest", string(data)}, " ")
}
