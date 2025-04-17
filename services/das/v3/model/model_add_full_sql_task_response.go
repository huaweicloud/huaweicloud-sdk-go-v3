package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddFullSqlTaskResponse Response Object
type AddFullSqlTaskResponse struct {

	// 任务ID
	TaskId         *int64 `json:"task_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o AddFullSqlTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddFullSqlTaskResponse struct{}"
	}

	return strings.Join([]string{"AddFullSqlTaskResponse", string(data)}, " ")
}
