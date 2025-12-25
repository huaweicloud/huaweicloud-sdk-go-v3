package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCollectorConnectionRequest Request Object
type UpdateCollectorConnectionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 连接 ID
	ConnectionId string `json:"connection_id"`

	Body *UpdateConnectionDto `json:"body,omitempty"`
}

func (o UpdateCollectorConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCollectorConnectionRequest struct{}"
	}

	return strings.Join([]string{"UpdateCollectorConnectionRequest", string(data)}, " ")
}
