package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowMigrationTaskStatsRequest struct {
	// 任务ID。

	TaskId string `json:"task_id"`
}

func (o ShowMigrationTaskStatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrationTaskStatsRequest struct{}"
	}

	return strings.Join([]string{"ShowMigrationTaskStatsRequest", string(data)}, " ")
}
