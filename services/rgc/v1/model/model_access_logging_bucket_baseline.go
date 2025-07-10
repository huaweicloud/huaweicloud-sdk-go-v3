package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessLoggingBucketBaseline 访问日志基础设置。
type AccessLoggingBucketBaseline struct {

	// 桶保留天数。
	RetentionDays int32 `json:"retention_days"`

	// 是否允许多AZ存储。
	EnableMultiAz *bool `json:"enable_multi_az,omitempty"`
}

func (o AccessLoggingBucketBaseline) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessLoggingBucketBaseline struct{}"
	}

	return strings.Join([]string{"AccessLoggingBucketBaseline", string(data)}, " ")
}
