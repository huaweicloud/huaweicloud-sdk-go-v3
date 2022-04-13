package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateConnectionRequest struct {
	// 连接名称.

	ConnectionName string `json:"connection_name"`

	Body *ConnectionInfo `json:"body,omitempty"`
}

func (o UpdateConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectionRequest struct{}"
	}

	return strings.Join([]string{"UpdateConnectionRequest", string(data)}, " ")
}
