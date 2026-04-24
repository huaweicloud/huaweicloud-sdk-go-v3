package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportSqlDataReq 导出流量回放任务SQL文件请求体
type ExportSqlDataReq struct {

	// 导出的sql文件类型。取值范围： - abnormal_sql ：异常sql列表 - abnormal_sql_detail ：异常sql详情 - slow_sql ：慢sql列表 - slow_sql_detail ： 慢sql详情
	FileType string `json:"file_type"`

	// 导出的字段名。file_type为时为error_sql_detail、slow_sql_detail必填。取值范围： -id ：自增ID -gmtCreate ：创建时间 -gmtModified ：修改时间 -shardId ：分片ID -startTimestamp ：源库SQL开始执行时间戳 -executeTimePartition ：回放时间分区 -schemaName ：库名 -queryType ：SQL类型 -sqlStatement ：SQL内容 -sqlTemplate ：SQL模板 -errorInfo ：异常信息 -targetType ：目标库类型
	FieldNames *[]string `json:"field_names,omitempty"`

	// 开始时间
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`
}

func (o ExportSqlDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportSqlDataReq struct{}"
	}

	return strings.Join([]string{"ExportSqlDataReq", string(data)}, " ")
}
