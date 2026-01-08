package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportDesktopListNewResponse Response Object
type ExportDesktopListNewResponse struct {

	// 任务id。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportDesktopListNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDesktopListNewResponse struct{}"
	}

	return strings.Join([]string{"ExportDesktopListNewResponse", string(data)}, " ")
}
