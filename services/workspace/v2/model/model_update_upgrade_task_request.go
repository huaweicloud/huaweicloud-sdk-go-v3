package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUpgradeTaskRequest Request Object
type UpdateUpgradeTaskRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`

	Body *UpdateScheduledUpgradeTaskRequestBody `json:"body,omitempty"`
}

func (o UpdateUpgradeTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUpgradeTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateUpgradeTaskRequest", string(data)}, " ")
}
