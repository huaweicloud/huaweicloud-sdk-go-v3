package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoggingBucketBaseline 日志基础设置。
type LoggingBucketBaseline struct {

	// 桶保留天数。
	RetentionDays int32 `json:"retention_days"`

	// 是否允许多AZ存储。
	EnableMultiAz *bool `json:"enable_multi_az,omitempty"`
}

func (o LoggingBucketBaseline) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoggingBucketBaseline struct{}"
	}

	return strings.Join([]string{"LoggingBucketBaseline", string(data)}, " ")
}
