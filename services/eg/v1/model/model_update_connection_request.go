package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConnectionRequest Request Object
type UpdateConnectionRequest struct {

	// 指定查询的目标连接ID
	ConnectionId string `json:"connection_id"`

	Body *ConnectionUpdateReq `json:"body,omitempty"`
}

func (o UpdateConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectionRequest struct{}"
	}

	return strings.Join([]string{"UpdateConnectionRequest", string(data)}, " ")
}
