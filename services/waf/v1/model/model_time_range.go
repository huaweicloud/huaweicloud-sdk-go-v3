package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TimeRange 时间范围
type TimeRange struct {

	// 开始时间戳（毫秒）
	From int64 `json:"from"`

	// 结束时间戳（毫秒）
	To int64 `json:"to"`
}

func (o TimeRange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeRange struct{}"
	}

	return strings.Join([]string{"TimeRange", string(data)}, " ")
}
