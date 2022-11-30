package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowTenantMetricRequest struct {

	// 时间段，单位为分钟
	Period *string `json:"period,omitempty"`

	// 开始时间，精确到ms的时间戳
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，精确到ms的时间戳
	EndTime *string `json:"end_time,omitempty"`

	// 指标类型，为空或不在取值范围内时，查询所有指标。取值范围：totalCount 调用次数；errorCount 错误次数； averageDuration 运行时间；running 运行中个数；rejectCount  拒绝个数。
	MetricType *string `json:"metric_type,omitempty"`
}

func (o ShowTenantMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantMetricRequest struct{}"
	}

	return strings.Join([]string{"ShowTenantMetricRequest", string(data)}, " ")
}
