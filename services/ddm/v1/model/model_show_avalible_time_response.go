package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvalibleTimeResponse Response Object
type ShowAvalibleTimeResponse struct {

	// 可恢复时间点。
	RestorableTimeIntervals *[]RestoreTimeInterval `json:"restorable_time_intervals,omitempty"`
	HttpStatusCode          int                    `json:"-"`
}

func (o ShowAvalibleTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvalibleTimeResponse struct{}"
	}

	return strings.Join([]string{"ShowAvalibleTimeResponse", string(data)}, " ")
}
