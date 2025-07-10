package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopResponse Response Object
type CreateDesktopResponse struct {

	// 创建云桌面总任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDesktopResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopResponse struct{}"
	}

	return strings.Join([]string{"CreateDesktopResponse", string(data)}, " ")
}
