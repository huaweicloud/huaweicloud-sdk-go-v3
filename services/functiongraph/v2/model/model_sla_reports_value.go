package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlaReportsValue struct {
	// 时间戳

	Timestamp *int32 `json:"timestamp,omitempty"`
	// 值

	Value *int32 `json:"value,omitempty"`
}

func (o SlaReportsValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlaReportsValue struct{}"
	}

	return strings.Join([]string{"SlaReportsValue", string(data)}, " ")
}
