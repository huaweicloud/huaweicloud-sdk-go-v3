package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDesktopPoolVolumesResponse Response Object
type AddDesktopPoolVolumesResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddDesktopPoolVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDesktopPoolVolumesResponse struct{}"
	}

	return strings.Join([]string{"AddDesktopPoolVolumesResponse", string(data)}, " ")
}
