package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeDesktopToImageResponse Response Object
type ChangeDesktopToImageResponse struct {

	// 桌面转镜像任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeDesktopToImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDesktopToImageResponse struct{}"
	}

	return strings.Join([]string{"ChangeDesktopToImageResponse", string(data)}, " ")
}
