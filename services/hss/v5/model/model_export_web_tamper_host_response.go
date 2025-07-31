package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportWebTamperHostResponse Response Object
type ExportWebTamperHostResponse struct {

	// 导出任务ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportWebTamperHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportWebTamperHostResponse struct{}"
	}

	return strings.Join([]string{"ExportWebTamperHostResponse", string(data)}, " ")
}
