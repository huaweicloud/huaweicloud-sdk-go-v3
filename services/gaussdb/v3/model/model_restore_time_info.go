package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreTimeInfo struct {

	// 可恢复时间段的起始时间点，UNIX时间戳格式，单位是毫秒，时区是UTC。
	StartTime *int64 `json:"start_time,omitempty"`

	// 可恢复时间段的结束时间点，UNIX时间戳格式，单位是毫秒，时区是UTC。
	EndTime *int64 `json:"end_time,omitempty"`
}

func (o RestoreTimeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreTimeInfo struct{}"
	}

	return strings.Join([]string{"RestoreTimeInfo", string(data)}, " ")
}
