package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowProgressDetailResponse struct {
	// 任务状态,0表示检查中，1表示检查失败，2表示检查成功，3表示任务中止

	TaskStatus *int32 `json:"task_status,omitempty"`

	Progress       *ProgressDetailV2 `json:"progress,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowProgressDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProgressDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowProgressDetailResponse", string(data)}, " ")
}
