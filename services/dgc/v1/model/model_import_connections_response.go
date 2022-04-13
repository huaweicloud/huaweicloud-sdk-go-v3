package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ImportConnectionsResponse struct {
	// 任务id

	TaskId         *string `json:"taskId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ImportConnectionsResponse", string(data)}, " ")
}
