package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUserConnectionNewResponse Response Object
type ExportUserConnectionNewResponse struct {

	// 导出任务id。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportUserConnectionNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserConnectionNewResponse struct{}"
	}

	return strings.Join([]string{"ExportUserConnectionNewResponse", string(data)}, " ")
}
