package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSharedInfoNewRequest Request Object
type UpdateSharedInfoNewRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *UpdateSharedInfoNewRequestBody `json:"body,omitempty"`
}

func (o UpdateSharedInfoNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSharedInfoNewRequest struct{}"
	}

	return strings.Join([]string{"UpdateSharedInfoNewRequest", string(data)}, " ")
}
