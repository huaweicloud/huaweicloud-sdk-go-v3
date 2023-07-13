package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesktopVolumesResponse Response Object
type DeleteDesktopVolumesResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDesktopVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesktopVolumesResponse struct{}"
	}

	return strings.Join([]string{"DeleteDesktopVolumesResponse", string(data)}, " ")
}
