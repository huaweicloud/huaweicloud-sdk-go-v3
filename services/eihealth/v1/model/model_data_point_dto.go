package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 监控数据列表
type DataPointDto struct {

	// 时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 数据单位
	Unit *string `json:"unit,omitempty"`

	// 统计结果
	Value float32 `json:"value,omitempty"`
}

func (o DataPointDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataPointDto struct{}"
	}

	return strings.Join([]string{"DataPointDto", string(data)}, " ")
}
