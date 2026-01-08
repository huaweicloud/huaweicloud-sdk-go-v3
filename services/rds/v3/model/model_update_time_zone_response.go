package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTimeZoneResponse Response Object
type UpdateTimeZoneResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTimeZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTimeZoneResponse struct{}"
	}

	return strings.Join([]string{"UpdateTimeZoneResponse", string(data)}, " ")
}
