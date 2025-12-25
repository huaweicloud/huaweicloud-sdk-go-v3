package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCollectorConnectionRequest Request Object
type ShowCollectorConnectionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 连接 ID
	ConnectionId string `json:"connection_id"`
}

func (o ShowCollectorConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCollectorConnectionRequest struct{}"
	}

	return strings.Join([]string{"ShowCollectorConnectionRequest", string(data)}, " ")
}
