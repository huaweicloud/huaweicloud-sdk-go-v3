package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDesktopVolumesResponse Response Object
type AddDesktopVolumesResponse struct {

	// 增加磁盘任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddDesktopVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDesktopVolumesResponse struct{}"
	}

	return strings.Join([]string{"AddDesktopVolumesResponse", string(data)}, " ")
}
