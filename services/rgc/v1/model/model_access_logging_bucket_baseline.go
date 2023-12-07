package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessLoggingBucketBaseline Access日志基础设置。
type AccessLoggingBucketBaseline struct {

	// 桶保留天数。
	RetentionDays int32 `json:"retention_days"`
}

func (o AccessLoggingBucketBaseline) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessLoggingBucketBaseline struct{}"
	}

	return strings.Join([]string{"AccessLoggingBucketBaseline", string(data)}, " ")
}
