package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebuildDesktopPoolResponse Response Object
type RebuildDesktopPoolResponse struct {

	// 重建系统盘总任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RebuildDesktopPoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebuildDesktopPoolResponse struct{}"
	}

	return strings.Join([]string{"RebuildDesktopPoolResponse", string(data)}, " ")
}
