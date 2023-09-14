package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScheduleRequest Request Object
type CreateScheduleRequest struct {

	// 边缘节点ID
	EdgeNodeId string `json:"edge_node_id"`

	Body *CreateScheduleReqDto `json:"body,omitempty"`
}

func (o CreateScheduleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScheduleRequest struct{}"
	}

	return strings.Join([]string{"CreateScheduleRequest", string(data)}, " ")
}
