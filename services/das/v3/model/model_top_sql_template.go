package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopSqlTemplate struct {

	// SQL模板。
	SqlTemplate string `json:"sql_template"`

	// SQL样本。
	SqlSampleString string `json:"sql_sample_string"`

	// SQL操作类型。
	SqlType string `json:"sql_type"`

	// 数据库名称。
	DbName string `json:"db_name"`

	// 总执行次数。
	ExecuteNum int64 `json:"execute_num"`

	// 总耗时(ms)。
	TotalCost float64 `json:"total_cost"`

	// 平均耗时(ms)。
	AvgCost float64 `json:"avg_cost"`

	// 平均返回行数。
	AvgRowsSent float64 `json:"avg_rows_sent"`

	// 平均影响行数。
	AvgRowsAffected float64 `json:"avg_rows_affected"`

	// 平均锁等待耗时(ms)。
	AvgLockTime float64 `json:"avg_lock_time"`

	// 总扫描行数。
	TotalRowsExamined float64 `json:"total_rows_examined"`

	// 平均扫描行数。
	AvgRowsExamined float64 `json:"avg_rows_examined"`

	// 总耗时占比。
	TotalCostRatio string `json:"total_cost_ratio"`

	// 扫描行数占比。
	TotalExaminedRatio string `json:"total_examined_ratio"`

	// 执行次数占比。
	ExecuteNumRatio string `json:"execute_num_ratio"`
}

func (o TopSqlTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopSqlTemplate struct{}"
	}

	return strings.Join([]string{"TopSqlTemplate", string(data)}, " ")
}
