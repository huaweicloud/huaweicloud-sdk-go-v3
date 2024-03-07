package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteScheduledTasksRequest Request Object
type BatchDeleteScheduledTasksRequest struct {
	Body *DeleteScheduledTasksReq `json:"body,omitempty"`
}

func (o BatchDeleteScheduledTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScheduledTasksRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteScheduledTasksRequest", string(data)}, " ")
}
