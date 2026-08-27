package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportDesktopVersionListResponse Response Object
type ExportDesktopVersionListResponse struct {

	// 任务id。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportDesktopVersionListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDesktopVersionListResponse struct{}"
	}

	return strings.Join([]string{"ExportDesktopVersionListResponse", string(data)}, " ")
}
