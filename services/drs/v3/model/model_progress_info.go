package model

import (
	"encoding/json"

	"strings"
)

// 迁移进度信息体
type ProgressInfo struct {
	// 完成进度

	Completed *string `json:"completed,omitempty"`
	// 预计剩余时间

	RemainingTime *string `json:"remaining_time,omitempty"`
}

func (o ProgressInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ProgressInfo struct{}"
	}

	return strings.Join([]string{"ProgressInfo", string(data)}, " ")
}
