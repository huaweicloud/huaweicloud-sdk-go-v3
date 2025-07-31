package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportVulHandleHistoryResponse Response Object
type ExportVulHandleHistoryResponse struct {

	// 导出任务ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportVulHandleHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportVulHandleHistoryResponse struct{}"
	}

	return strings.Join([]string{"ExportVulHandleHistoryResponse", string(data)}, " ")
}
