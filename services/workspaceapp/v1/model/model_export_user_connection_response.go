package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportUserConnectionResponse Response Object
type ExportUserConnectionResponse struct {

	// 导出任务id。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportUserConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportUserConnectionResponse struct{}"
	}

	return strings.Join([]string{"ExportUserConnectionResponse", string(data)}, " ")
}
