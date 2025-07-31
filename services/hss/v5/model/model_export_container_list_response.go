package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportContainerListResponse Response Object
type ExportContainerListResponse struct {

	// 导出任务ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportContainerListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportContainerListResponse struct{}"
	}

	return strings.Join([]string{"ExportContainerListResponse", string(data)}, " ")
}
