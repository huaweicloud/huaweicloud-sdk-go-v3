package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TplCmp SQL模板对比项
type TplCmp struct {

	// SQL模板ID
	TemplateId *string `json:"template_id,omitempty"`

	// SQL模板
	SqlTemplate *string `json:"sql_template,omitempty"`

	// SQL样例
	SqlSampleString *string `json:"sql_sample_string,omitempty"`

	// 数据库列表
	DbNames *[]string `json:"db_names,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// SQL类型
	SqlType *string `json:"sql_type,omitempty"`

	// 执行次数
	ExecuteNum *int64 `json:"execute_num,omitempty"`

	// 总执行耗时 ms
	TotalCost *float64 `json:"total_cost,omitempty"`

	// 平均执行耗时 ms
	AvgCost *float64 `json:"avg_cost,omitempty"`

	// 最大执行耗时 ms
	MaxCost *float64 `json:"max_cost,omitempty"`

	// 平均返回行数
	AvgRowsSent *float64 `json:"avg_rows_sent,omitempty"`

	// 最大返回行数
	MaxRowsSent *float64 `json:"max_rows_sent,omitempty"`

	// 平均影响行数
	AvgRowsAffected *float64 `json:"avg_rows_affected,omitempty"`

	// 最大影响行数
	MaxRowsAffected *float64 `json:"max_rows_affected,omitempty"`

	// 平均锁等待时间
	AvgLockTime *float64 `json:"avg_lock_time,omitempty"`

	// 最大锁等待时间
	MaxLockTime *float64 `json:"max_lock_time,omitempty"`

	// 总扫描行数
	TotalRowsExamined *float64 `json:"total_rows_examined,omitempty"`

	// 平均扫描行数
	AvgRowsExamined *float64 `json:"avg_rows_examined,omitempty"`

	// 最大扫描行数
	MaxRowsExamined *float64 `json:"max_rows_examined,omitempty"`

	// 执行耗时占比
	TotalCostRatio *string `json:"total_cost_ratio,omitempty"`

	// 扫描行数占比
	TotalExaminedRatio *string `json:"total_examined_ratio,omitempty"`

	// 执行次数占比
	ExecuteNumRatio *string `json:"execute_num_ratio,omitempty"`
}

func (o TplCmp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TplCmp struct{}"
	}

	return strings.Join([]string{"TplCmp", string(data)}, " ")
}
