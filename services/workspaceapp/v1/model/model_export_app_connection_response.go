package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportAppConnectionResponse Response Object
type ExportAppConnectionResponse struct {

	// 导出任务id。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportAppConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportAppConnectionResponse struct{}"
	}

	return strings.Join([]string{"ExportAppConnectionResponse", string(data)}, " ")
}
