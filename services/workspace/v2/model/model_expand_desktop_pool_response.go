package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandDesktopPoolResponse Response Object
type ExpandDesktopPoolResponse struct {

	// 创建云桌面总任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandDesktopPoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandDesktopPoolResponse struct{}"
	}

	return strings.Join([]string{"ExpandDesktopPoolResponse", string(data)}, " ")
}
