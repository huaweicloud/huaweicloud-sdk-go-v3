package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkflowScheduleListResponse Response Object
type ShowWorkflowScheduleListResponse struct {

	// **参数解释**：工作流定时调度列表
	Schedules      *[]WorkflowScheduleResp `json:"schedules,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowWorkflowScheduleListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowScheduleListResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkflowScheduleListResponse", string(data)}, " ")
}
