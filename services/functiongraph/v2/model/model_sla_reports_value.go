package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlaReportsValue struct {

	// 时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 值
	Value *float64 `json:"value,omitempty"`
}

func (o SlaReportsValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlaReportsValue struct{}"
	}

	return strings.Join([]string{"SlaReportsValue", string(data)}, " ")
}
