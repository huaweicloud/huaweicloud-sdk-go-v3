package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeDesktopResponse Response Object
type ResizeDesktopResponse struct {

	// 按需桌面变更规格返回的任务信息（jobs字段后续会下线，请使用job_id字段）。
	Jobs *[]ResizeDesktopJobResponse `json:"jobs,omitempty"`

	// 变更规格任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeDesktopResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopResponse struct{}"
	}

	return strings.Join([]string{"ResizeDesktopResponse", string(data)}, " ")
}
