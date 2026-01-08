package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUsersNewResponse Response Object
type ExportUsersNewResponse struct {

	// 导出任务id。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportUsersNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUsersNewResponse struct{}"
	}

	return strings.Join([]string{"ExportUsersNewResponse", string(data)}, " ")
}
