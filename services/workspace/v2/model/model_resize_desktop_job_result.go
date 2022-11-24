package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 变更桌面规格Job响应。
type ResizeDesktopJobResult struct {

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`
}

func (o ResizeDesktopJobResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopJobResult struct{}"
	}

	return strings.Join([]string{"ResizeDesktopJobResult", string(data)}, " ")
}
