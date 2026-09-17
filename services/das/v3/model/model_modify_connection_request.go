package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyConnectionRequest Request Object
type ModifyConnectionRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *ModifyConnectionRequestBody `json:"body,omitempty"`
}

func (o ModifyConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyConnectionRequest struct{}"
	}

	return strings.Join([]string{"ModifyConnectionRequest", string(data)}, " ")
}
