package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 删除迁移任务请求体
type DeleteMigrateTaskRequest struct {
	// 删除的迁移任务ID列表。

	TaskIdList []string `json:"task_id_list"`
}

func (o DeleteMigrateTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMigrateTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteMigrateTaskRequest", string(data)}, " ")
}
