package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRebuildDesktopsSystemDiskResponse Response Object
type BatchRebuildDesktopsSystemDiskResponse struct {

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 重建系统盘总任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchRebuildDesktopsSystemDiskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebuildDesktopsSystemDiskResponse struct{}"
	}

	return strings.Join([]string{"BatchRebuildDesktopsSystemDiskResponse", string(data)}, " ")
}
