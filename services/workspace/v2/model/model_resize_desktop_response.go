package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ResizeDesktopResponse struct {

	// 按需桌面变更规格返回的任务信息。
	Jobs           *[]ResizeDesktopJobResult `json:"jobs,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ResizeDesktopResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeDesktopResponse struct{}"
	}

	return strings.Join([]string{"ResizeDesktopResponse", string(data)}, " ")
}
