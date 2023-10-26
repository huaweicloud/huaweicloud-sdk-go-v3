package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckMigrationConnectivityRequest Request Object
type CheckMigrationConnectivityRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`

	Body *RedisConnectionParam `json:"body,omitempty"`
}

func (o CheckMigrationConnectivityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckMigrationConnectivityRequest struct{}"
	}

	return strings.Join([]string{"CheckMigrationConnectivityRequest", string(data)}, " ")
}
