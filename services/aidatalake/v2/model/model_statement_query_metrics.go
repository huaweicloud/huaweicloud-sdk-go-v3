package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatementQueryMetrics **参数解释**：语句监控详情。
type StatementQueryMetrics struct {

	// **参数解释**：会话ID。 **取值范围**：长度为1~36个字符，支持大小写英文字母、数字、连字符。
	SessionId *string `json:"session_id,omitempty"`

	// **参数解释**：语句ID。 **取值范围**：不涉及。
	StatementId *string `json:"statement_id,omitempty"`

	// **参数解释**：语句开始时间。 **取值范围**：不涉及。
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**：语句运行时长（单位：ms）。 **取值范围**：不涉及。
	Duration *int64 `json:"duration,omitempty"`

	// **参数解释**：语句执行时的Schema名称。 **取值范围**：不涉及。
	SchemaName *string `json:"schema_name,omitempty"`

	// **参数解释**：查询语句。 **取值范围**：不涉及。
	Query *string `json:"query,omitempty"`

	// **参数解释**：语句计划信息。 **取值范围**：不涉及。
	QueryPlan *string `json:"query_plan,omitempty"`

	// **参数解释**：预计语句运行时长（单位：ms）。 **取值范围**：不涉及。
	EstimateTotalTime *int64 `json:"estimate_total_time,omitempty"`

	// **参数解释**：预计语句剩余时长（单位：ms）。 **取值范围**：不涉及。
	EstimateLeftTime *int64 `json:"estimate_left_time,omitempty"`

	// **参数解释**：cn启动时长（单位：ms）。 **取值范围**：不涉及。
	StartupCoordinatorDuration *int64 `json:"startup_coordinator_duration,omitempty"`

	// **参数解释**：创建会话时长（单位：ms）。 **取值范围**：不涉及。
	CreateSessionDuration *int64 `json:"create_session_duration,omitempty"`

	// **参数解释**：语句加入任务队列时长（单位：ms）。 **取值范围**：不涉及。
	EnqueueStatementDuration *int64 `json:"enqueue_statement_duration,omitempty"`

	// **参数解释**：语句排队时长（单位：ms）。 **取值范围**：不涉及。
	QueueDuration *int64 `json:"queue_duration,omitempty"`

	// **参数解释**：语句写结果集时长（单位：ms）。 **取值范围**：不涉及。
	WriteResultsetDuration *int64 `json:"write_resultset_duration,omitempty"`

	// **参数解释**：执行前解析时间（单位：ms）。 **取值范围**：不涉及。
	ParseTime *int64 `json:"parse_time,omitempty"`

	// **参数解释**：预计使用内存（单位：MB）。 **取值范围**：不涉及。
	EstimateMemory *int32 `json:"estimate_memory,omitempty"`

	// **参数解释**：语句状态，WLM状态。 **取值范围**：不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：语句报错信息。 **取值范围**：不涉及。
	ErrorInfo *string `json:"error_info,omitempty"`

	// **参数解释**：DN时长偏斜百分比（单位：%）。 **取值范围**：不涉及。
	DnTimeSkewPercent *int32 `json:"dn_time_skew_percent,omitempty"`

	// **参数解释**：DN最大峰值内存（单位：MB）。 **取值范围**：不涉及。
	DnMaxPeakMemory *int32 `json:"dn_max_peak_memory,omitempty"`

	// **参数解释**：DN平均峰值内存（单位：MB）。 **取值范围**：不涉及。
	DnAveragePeakMemory *int32 `json:"dn_average_peak_memory,omitempty"`

	// **参数解释**：dn下盘倾斜率（单位：%）。 **取值范围**：不涉及。
	DnSpillSkewPercent *int32 `json:"dn_spill_skew_percent,omitempty"`

	// **参数解释**：dn落盘总大小（单位：MB）。 **取值范围**：不涉及。
	DnWriteDiskTotalSize *int64 `json:"dn_write_disk_total_size,omitempty"`

	// **参数解释**：dn溢写OBS总大小（单位：MB）。 **取值范围**：不涉及。
	DnSpillObsTotalSize *int64 `json:"dn_spill_obs_total_size,omitempty"`

	// **参数解释**：cn溢写OBS大小（单位：MB）。 **取值范围**：不涉及。
	CnSpillObsSize *int64 `json:"cn_spill_obs_size,omitempty"`

	// **参数解释**：磁盘缓存命中率。 **取值范围**：不涉及。
	DiskCacheHitRatio float32 `json:"disk_cache_hit_ratio,omitempty"`

	// **参数解释**：OBS IO请求平均延迟（单位：us）。 **取值范围**：不涉及。
	ObsIoReqAvgLatency *int64 `json:"obs_io_req_avg_latency,omitempty"`

	// **参数解释**：OBS IO请求延迟大于1秒的次数。 **取值范围**：不涉及。
	ObsIoReqLatencyGt1s *int64 `json:"obs_io_req_latency_gt_1s,omitempty"`

	// **参数解释**：OBS IO请求延迟大于10秒的次数。 **取值范围**：不涉及。
	ObsIoReqLatencyGt10s *int64 `json:"obs_io_req_latency_gt_10s,omitempty"`
}

func (o StatementQueryMetrics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatementQueryMetrics struct{}"
	}

	return strings.Join([]string{"StatementQueryMetrics", string(data)}, " ")
}
