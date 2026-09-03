package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTuningRequest Request Object
type ExecuteTuningRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *ExecuteTuningRequestBody `json:"body,omitempty"`
}

func (o ExecuteTuningRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTuningRequest struct{}"
	}

	return strings.Join([]string{"ExecuteTuningRequest", string(data)}, " ")
}
