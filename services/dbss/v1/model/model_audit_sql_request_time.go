package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuditSqlRequestTime 查询时间范围
type AuditSqlRequestTime struct {

	// 时间范围。和start_time，end_time不能同时使用，同时传该参数优先级更高。 - HALF_HOUR：半小时 - HOUR：1小时 - THREE_HOUR：3小时 - TWELVE_HOUR：12小时 - DAY：1天 - WEEK：一周 - MONTH：一个月
	TimeRange *string `json:"time_range,omitempty"`

	// 开始时间，必须和end_time成对出现。格式必须为yyyy-MM-dd HH:mm:ss。UTC时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，必须和start_time成对出现。格式必须为yyyy-MM-dd HH:mm:ss。UTC时间
	EndTime *string `json:"end_time,omitempty"`
}

func (o AuditSqlRequestTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditSqlRequestTime struct{}"
	}

	return strings.Join([]string{"AuditSqlRequestTime", string(data)}, " ")
}
