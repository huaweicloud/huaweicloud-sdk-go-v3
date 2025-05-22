package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecycleBinPolicy struct {

	// 允许进入回收站的最小创建时间，不足此时长则删除时不满足进入回收站的条件。
	RecycleThresholdHour *int32 `json:"recycle_threshold_hour,omitempty"`

	// 进入回收站的最大保留时长。
	RetentionHour *int32 `json:"retention_hour,omitempty"`
}

func (o RecycleBinPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleBinPolicy struct{}"
	}

	return strings.Join([]string{"RecycleBinPolicy", string(data)}, " ")
}
