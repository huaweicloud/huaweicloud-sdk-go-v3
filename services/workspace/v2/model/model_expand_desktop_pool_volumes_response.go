package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandDesktopPoolVolumesResponse Response Object
type ExpandDesktopPoolVolumesResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandDesktopPoolVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandDesktopPoolVolumesResponse struct{}"
	}

	return strings.Join([]string{"ExpandDesktopPoolVolumesResponse", string(data)}, " ")
}
