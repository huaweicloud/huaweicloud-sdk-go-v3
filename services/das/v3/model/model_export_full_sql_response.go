package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportFullSqlResponse Response Object
type ExportFullSqlResponse struct {

	// 任务ID
	TaskId         *int64 `json:"task_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ExportFullSqlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFullSqlResponse struct{}"
	}

	return strings.Join([]string{"ExportFullSqlResponse", string(data)}, " ")
}
