package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProgressCompleteInfo 全量迁移进度。
type ProgressCompleteInfo struct {

	// 完成进度。
	Completed *string `json:"completed,omitempty"`

	// 预计剩余时间。
	RemainingTime *string `json:"remaining_time,omitempty"`
}

func (o ProgressCompleteInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProgressCompleteInfo struct{}"
	}

	return strings.Join([]string{"ProgressCompleteInfo", string(data)}, " ")
}
