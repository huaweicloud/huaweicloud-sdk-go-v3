package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateConnectionRequest struct {

	// 连接名称.
	ConnectionName string `json:"connection_name" xml:"connection_name"`

	Body *ConnectionInfo `json:"body,omitempty" xml:"body"`
}

func (o UpdateConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectionRequest struct{}"
	}

	return strings.Join([]string{"UpdateConnectionRequest", string(data)}, " ")
}
