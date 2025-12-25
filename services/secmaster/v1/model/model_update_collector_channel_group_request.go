package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCollectorChannelGroupRequest Request Object
type UpdateCollectorChannelGroupRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 组ID
	GroupId string `json:"group_id"`

	Body *GroupDto `json:"body,omitempty"`
}

func (o UpdateCollectorChannelGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCollectorChannelGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateCollectorChannelGroupRequest", string(data)}, " ")
}
