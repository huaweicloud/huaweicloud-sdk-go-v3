package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandDesktopVolumeResponse Response Object
type ExpandDesktopVolumeResponse struct {

	// 扩容磁盘任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandDesktopVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandDesktopVolumeResponse struct{}"
	}

	return strings.Join([]string{"ExpandDesktopVolumeResponse", string(data)}, " ")
}
