package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCollectorChannelGroupRequest Request Object
type ListCollectorChannelGroupRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 名称
	Name *string `json:"name,omitempty"`
}

func (o ListCollectorChannelGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCollectorChannelGroupRequest struct{}"
	}

	return strings.Join([]string{"ListCollectorChannelGroupRequest", string(data)}, " ")
}
