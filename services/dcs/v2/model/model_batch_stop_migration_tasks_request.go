package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStopMigrationTasksRequest Request Object
type BatchStopMigrationTasksRequest struct {
	Body *BatchStopMigrationTasksBody `json:"body,omitempty"`
}

func (o BatchStopMigrationTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopMigrationTasksRequest struct{}"
	}

	return strings.Join([]string{"BatchStopMigrationTasksRequest", string(data)}, " ")
}
