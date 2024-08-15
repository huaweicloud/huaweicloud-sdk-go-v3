package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRestartOnlineMigrationTasksRequest Request Object
type BatchRestartOnlineMigrationTasksRequest struct {
	Body *BatchRestartOnlineMigrationTasksBody `json:"body,omitempty"`
}

func (o BatchRestartOnlineMigrationTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRestartOnlineMigrationTasksRequest struct{}"
	}

	return strings.Join([]string{"BatchRestartOnlineMigrationTasksRequest", string(data)}, " ")
}
