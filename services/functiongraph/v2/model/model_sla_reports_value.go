package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlaReportsValue struct {

	// 时间戳
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp"`

	// 值
	Value *float64 `json:"value,omitempty" xml:"value"`
}

func (o SlaReportsValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlaReportsValue struct{}"
	}

	return strings.Join([]string{"SlaReportsValue", string(data)}, " ")
}
