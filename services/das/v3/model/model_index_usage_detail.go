package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndexUsageDetail IndexUsageDetail
type IndexUsageDetail struct {

	// 表名称
	TableName *string `json:"table_name,omitempty"`

	// 索引名称
	IndexName *string `json:"index_name,omitempty"`

	// 索引类型描述
	IxTypeDesc *string `json:"ix_type_desc,omitempty"`

	// 碎片率
	FragmentationPercentage *float64 `json:"fragmentation_percentage,omitempty"`

	// 索引占用的空间大小(MB)
	IndexSizeMb *float64 `json:"index_size_mb,omitempty"`

	// 维护操作
	MaintenanceOperation *string `json:"maintenance_operation,omitempty"`

	// 索引占用的空间页数
	PageCount *int64 `json:"page_count,omitempty"`

	// 通过用户查询执行的搜索次数
	IxSeekCount *int64 `json:"ix_seek_count,omitempty"`

	// 未使用索引的用户查询的扫描数
	IxScanCount *int64 `json:"ix_scan_count,omitempty"`

	// 由用户查询执行的书签查找次数
	IxKeyLookupCount *int64 `json:"ix_key_lookup_count,omitempty"`

	// 通过用户查询执行的更新次数
	IxUpdateCount *int64 `json:"ix_update_count,omitempty"`

	// 查找百分比
	SeekPercentage *float64 `json:"seek_percentage,omitempty"`

	// 扫描百分比
	ScanPercentage *float64 `json:"scan_percentage,omitempty"`

	// 书签查找百分比
	KeyLookupPercentage *float64 `json:"key_lookup_percentage,omitempty"`

	// 更新百分比
	UpdatePercentage *float64 `json:"update_percentage,omitempty"`

	// 索引是否是主键
	IsPrimaryKey *bool `json:"is_primary_key,omitempty"`

	// 索引是否被禁用
	IsDisabled *bool `json:"is_disabled,omitempty"`

	// 列
	ColumnList *string `json:"column_list,omitempty"`

	// 填充因子
	FillFactor *string `json:"fill_factor,omitempty"`

	// 创建时间
	CreateDate *int64 `json:"create_date,omitempty"`

	// 统计信息更新时间
	StatsLastUpdated *int64 `json:"stats_last_updated,omitempty"`
}

func (o IndexUsageDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndexUsageDetail struct{}"
	}

	return strings.Join([]string{"IndexUsageDetail", string(data)}, " ")
}
