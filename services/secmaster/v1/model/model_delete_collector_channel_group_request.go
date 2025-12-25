package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCollectorChannelGroupRequest Request Object
type DeleteCollectorChannelGroupRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 组ID
	GroupId string `json:"group_id"`
}

func (o DeleteCollectorChannelGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCollectorChannelGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteCollectorChannelGroupRequest", string(data)}, " ")
}
