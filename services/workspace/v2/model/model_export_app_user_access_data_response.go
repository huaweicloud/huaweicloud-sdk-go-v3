package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportAppUserAccessDataResponse Response Object
type ExportAppUserAccessDataResponse struct {

	// 导出任务id。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportAppUserAccessDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportAppUserAccessDataResponse struct{}"
	}

	return strings.Join([]string{"ExportAppUserAccessDataResponse", string(data)}, " ")
}
