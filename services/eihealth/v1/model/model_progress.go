package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Progress struct {

	// 整体进度
	Overall *float32 `json:"overall,omitempty"`

	// 预计结束时间，毫秒
	EstimatedFinishTime *int64 `json:"estimated_finish_time,omitempty"`
}

func (o Progress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Progress struct{}"
	}

	return strings.Join([]string{"Progress", string(data)}, " ")
}
