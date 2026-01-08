package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeDesktopPoolResponse Response Object
type ResizeDesktopPoolResponse struct {

	// 创建云桌面总任务id。
	JobId *string `json:"job_id,omitempty"`

	// 桌面变更规格返回的任务信息
	Jobs           *[]ResizeDesktopPoolJobResponse `json:"jobs,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ResizeDesktopPoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopPoolResponse struct{}"
	}

	return strings.Join([]string{"ResizeDesktopPoolResponse", string(data)}, " ")
}
