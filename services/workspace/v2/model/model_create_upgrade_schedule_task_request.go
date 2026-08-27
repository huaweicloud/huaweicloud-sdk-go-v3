package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUpgradeScheduleTaskRequest Request Object
type CreateUpgradeScheduleTaskRequest struct {
	Body *CreateScheduledUpgradeTaskRequestBody `json:"body,omitempty"`
}

func (o CreateUpgradeScheduleTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUpgradeScheduleTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateUpgradeScheduleTaskRequest", string(data)}, " ")
}
