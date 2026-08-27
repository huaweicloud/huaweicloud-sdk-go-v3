package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUserGroupsNewResponse Response Object
type ExportUserGroupsNewResponse struct {

	// 导出任务id。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportUserGroupsNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserGroupsNewResponse struct{}"
	}

	return strings.Join([]string{"ExportUserGroupsNewResponse", string(data)}, " ")
}
