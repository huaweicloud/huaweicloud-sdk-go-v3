package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteTasksRequest Request Object
type BatchDeleteTasksRequest struct {

	// 项目id
	ProjectUuid string `json:"project_uuid"`

	Body *DeleteTaskInfo `json:"body,omitempty"`
}

func (o BatchDeleteTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTasksRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteTasksRequest", string(data)}, " ")
}
