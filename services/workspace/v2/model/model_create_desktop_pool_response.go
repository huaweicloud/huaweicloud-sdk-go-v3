package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopPoolResponse Response Object
type CreateDesktopPoolResponse struct {

	// 创建云桌面总任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDesktopPoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopPoolResponse struct{}"
	}

	return strings.Join([]string{"CreateDesktopPoolResponse", string(data)}, " ")
}
