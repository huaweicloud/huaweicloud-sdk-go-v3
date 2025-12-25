package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCollectorConnectionRequest Request Object
type DeleteCollectorConnectionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 连接 ID
	ConnectionId string `json:"connection_id"`
}

func (o DeleteCollectorConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCollectorConnectionRequest struct{}"
	}

	return strings.Join([]string{"DeleteCollectorConnectionRequest", string(data)}, " ")
}
