package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSearchPathFlagRequest Request Object
type UpdateSearchPathFlagRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *UpdateSearchPathFlagRequestBody `json:"body,omitempty"`
}

func (o UpdateSearchPathFlagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSearchPathFlagRequest struct{}"
	}

	return strings.Join([]string{"UpdateSearchPathFlagRequest", string(data)}, " ")
}
