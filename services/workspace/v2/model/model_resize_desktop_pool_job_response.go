package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeDesktopPoolJobResponse 按需桌面变更规格返回的任务信息
type ResizeDesktopPoolJobResponse struct {

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`
}

func (o ResizeDesktopPoolJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopPoolJobResponse struct{}"
	}

	return strings.Join([]string{"ResizeDesktopPoolJobResponse", string(data)}, " ")
}
