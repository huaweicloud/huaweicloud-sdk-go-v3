package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUserGroupUsersNewResponse Response Object
type ExportUserGroupUsersNewResponse struct {

	// 导出任务id。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportUserGroupUsersNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserGroupUsersNewResponse struct{}"
	}

	return strings.Join([]string{"ExportUserGroupUsersNewResponse", string(data)}, " ")
}
