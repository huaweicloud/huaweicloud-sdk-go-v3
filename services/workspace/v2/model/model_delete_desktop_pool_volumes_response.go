package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDesktopPoolVolumesResponse Response Object
type DeleteDesktopPoolVolumesResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDesktopPoolVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDesktopPoolVolumesResponse struct{}"
	}

	return strings.Join([]string{"DeleteDesktopPoolVolumesResponse", string(data)}, " ")
}
