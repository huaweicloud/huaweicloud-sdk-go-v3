package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExecutionPlanRequest Request Object
type ShowExecutionPlanRequest struct {

	// 连接ID
	ConnectionId string `json:"connection_id"`

	Body *ShowExecutionPlanRequestBody `json:"body,omitempty"`
}

func (o ShowExecutionPlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExecutionPlanRequest struct{}"
	}

	return strings.Join([]string{"ShowExecutionPlanRequest", string(data)}, " ")
}
