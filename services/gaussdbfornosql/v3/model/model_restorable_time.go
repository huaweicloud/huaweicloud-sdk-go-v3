package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestorableTime 可恢复的时间段
type RestorableTime struct {

	// 可恢复时间段的开始时间点，UNIX时间戳格式，单位是毫秒，时区是UTC。
	StartTime int64 `json:"start_time"`

	// 可恢复时间段的结束时间点， UNIX时间戳格式，单位是毫秒，时区是UTC。
	EndTime int64 `json:"end_time"`
}

func (o RestorableTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestorableTime struct{}"
	}

	return strings.Join([]string{"RestorableTime", string(data)}, " ")
}
