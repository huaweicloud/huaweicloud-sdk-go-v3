package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoggingConfiguration Landing Zone日志配置。
type LoggingConfiguration struct {

	// 日志桶名称。
	LoggingBucketName *string `json:"logging_bucket_name,omitempty"`

	AccessLoggingBucket *AccessLoggingBucketBaseline `json:"access_logging_bucket,omitempty"`

	LoggingBucket *LoggingBucketBaseline `json:"logging_bucket,omitempty"`
}

func (o LoggingConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoggingConfiguration struct{}"
	}

	return strings.Join([]string{"LoggingConfiguration", string(data)}, " ")
}
