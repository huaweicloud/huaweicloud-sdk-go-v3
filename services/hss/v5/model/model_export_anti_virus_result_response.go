package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportAntiVirusResultResponse Response Object
type ExportAntiVirusResultResponse struct {

	// 导出任务ID
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportAntiVirusResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportAntiVirusResultResponse struct{}"
	}

	return strings.Join([]string{"ExportAntiVirusResultResponse", string(data)}, " ")
}
