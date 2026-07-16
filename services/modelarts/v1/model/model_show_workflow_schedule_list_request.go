package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkflowScheduleListRequest Request Object
type ShowWorkflowScheduleListRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`
}

func (o ShowWorkflowScheduleListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowScheduleListRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkflowScheduleListRequest", string(data)}, " ")
}
