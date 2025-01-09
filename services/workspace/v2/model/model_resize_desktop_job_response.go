package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeDesktopJobResponse 变更桌面规格Job响应。
type ResizeDesktopJobResponse struct {

	// 错误码，失败时返回。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`
}

func (o ResizeDesktopJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopJobResponse struct{}"
	}

	return strings.Join([]string{"ResizeDesktopJobResponse", string(data)}, " ")
}
