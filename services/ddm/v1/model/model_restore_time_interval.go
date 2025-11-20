package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreTimeInterval struct {

	// 开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o RestoreTimeInterval) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreTimeInterval struct{}"
	}

	return strings.Join([]string{"RestoreTimeInterval", string(data)}, " ")
}
