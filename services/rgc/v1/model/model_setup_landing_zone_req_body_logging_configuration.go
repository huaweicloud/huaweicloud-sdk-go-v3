package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetupLandingZoneReqBodyLoggingConfiguration 日志桶设置信息。
type SetupLandingZoneReqBodyLoggingConfiguration struct {
	LoggingBucket *LoggingBucketBaseline `json:"logging_bucket"`

	AccessLoggingBucket *AccessLoggingBucketBaseline `json:"access_logging_bucket"`
}

func (o SetupLandingZoneReqBodyLoggingConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetupLandingZoneReqBodyLoggingConfiguration struct{}"
	}

	return strings.Join([]string{"SetupLandingZoneReqBodyLoggingConfiguration", string(data)}, " ")
}
