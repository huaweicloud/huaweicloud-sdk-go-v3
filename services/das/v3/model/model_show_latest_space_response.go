package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLatestSpaceResponse Response Object
type ShowLatestSpaceResponse struct {

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// MySQL购买空间
	BuyStorageBytes *int64 `json:"buy_storage_bytes,omitempty"`

	// MySQL已使用空间
	UsedStorageBytes *int64 `json:"used_storage_bytes,omitempty"`

	// MySQL剩余空间
	LastStorageBytes *int64 `json:"last_storage_bytes,omitempty"`

	// MySQL使用空间占比
	BuyStoragePercent *float64 `json:"buy_storage_percent,omitempty"`

	// 数据空间
	DataUsageBytes *int64 `json:"data_usage_bytes,omitempty"`

	// 数据空间占比
	DataUsagePercent *float64 `json:"data_usage_percent,omitempty"`

	// binlog空间
	BinlogUsageBytes *int64 `json:"binlog_usage_bytes,omitempty"`

	// binlog空间占比
	BinlogUsagePercent *float64 `json:"binlog_usage_percent,omitempty"`

	// relayLog空间
	RelayLogUsageBytes *int64 `json:"relay_log_usage_bytes,omitempty"`

	// relayLog空间占比
	RelayLogUsagePercent *float64 `json:"relay_log_usage_percent,omitempty"`

	// auditLog空间
	AuditLogUsageBytes *int64 `json:"audit_log_usage_bytes,omitempty"`

	// auditLog空间占比
	AuditLogUsagePercent *float64 `json:"audit_log_usage_percent,omitempty"`

	// slowLog空间
	SlowLogUsageBytes *int64 `json:"slow_log_usage_bytes,omitempty"`

	// slowLog空间占比
	SlowLogUsagePercent *float64 `json:"slow_log_usage_percent,omitempty"`

	// 临时空间
	TempUsageBytes *int64 `json:"temp_usage_bytes,omitempty"`

	// 临时空间占比
	TempUsagePercent *float64 `json:"temp_usage_percent,omitempty"`

	// undoLog空间
	UndoLogUsageBytes *int64 `json:"undo_log_usage_bytes,omitempty"`

	// undoLog空间占比
	UndoLogUsagePercent *float64 `json:"undo_log_usage_percent,omitempty"`

	// 其他空间
	OtherUsageBytes *int64 `json:"other_usage_bytes,omitempty"`

	// 其他空间占比
	OtherUsagePercent *float64 `json:"other_usage_percent,omitempty"`

	// 是否文件过多
	TooManyFiles *bool `json:"too_many_files,omitempty"`

	// TaurusDB数据空间
	PageUsageBytes *float64 `json:"page_usage_bytes,omitempty"`

	// SQLServer总空间
	TotalSpace *int64 `json:"total_space,omitempty"`

	// SQLServer已使用空间
	TotalUsage *int64 `json:"total_usage,omitempty"`

	// SQLServer可用空间
	AvailSize *int64 `json:"avail_size,omitempty"`

	// SQLServer数据空间
	Data *float64 `json:"data,omitempty"`

	// SQLServer log空间
	Log *float64 `json:"log,omitempty"`

	// SQLServer runtime
	Runtime *float64 `json:"runtime,omitempty"`

	// SQLServer slow_log空间
	SlowLog *float64 `json:"slow_log,omitempty"`

	// SQLServer audit_log空间
	AuditLog *float64 `json:"audit_log,omitempty"`

	// SQLServer tempdb空间
	Tempdb *float64 `json:"tempdb,omitempty"`

	// SQLServer msdb空间
	Msdb *float64 `json:"msdb,omitempty"`

	// DDS磁盘使用量
	UsedSizeBytes *int64 `json:"used_size_bytes,omitempty"`

	// DDS磁盘总量
	TotalSizeBytes *int64 `json:"total_size_bytes,omitempty"`

	// DDS近一周日均增长
	AvgDailyGrowthBytes *float64 `json:"avg_daily_growth_bytes,omitempty"`

	// DDS预计可用天数
	EstimatedAvailableDays *int64 `json:"estimated_available_days,omitempty"`

	// DDS数据空间
	DataSizeBytes *int64 `json:"data_size_bytes,omitempty"`

	// DDS oplog空间
	OplogSizeBytes *int64 `json:"oplog_size_bytes,omitempty"`

	// DDS其他空间
	OtherSizeBytes *int64 `json:"other_size_bytes,omitempty"`

	// PostgreSQL wallog空间
	WalSize *float64 `json:"wal_size,omitempty"`

	// PostgreSQL数据空间
	DataSize *float64 `json:"data_size,omitempty"`

	// PostgreSQL auditlog空间
	PgauditLogSize *float64 `json:"pgaudit_log_size,omitempty"`

	// PostgreSQL临时空间
	PgsqlTmpSize   *float64 `json:"pgsql_tmp_size,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowLatestSpaceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLatestSpaceResponse struct{}"
	}

	return strings.Join([]string{"ShowLatestSpaceResponse", string(data)}, " ")
}
