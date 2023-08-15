package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProfileInfo 概要信息
type ProfileInfo struct {

	// 数据库类型
	DbType *string `json:"db_type,omitempty"`

	// 是否取消
	Cancel *bool `json:"cancel,omitempty"`

	// 表大小
	TableSize float32 `json:"table_size,omitempty"`

	// 数据库名
	DatabaseName *string `json:"database_name,omitempty"`

	// obs公共配置
	ObsCommonConfig *string `json:"obs_common_config,omitempty"`

	// 总行数
	TotalRowCount *string `json:"total_row_count,omitempty"`

	// 文件列表
	FieldsName *[]string `json:"fields_name,omitempty"`

	// 表名
	TableName *string `json:"table_name,omitempty"`

	// 样本
	Sample *string `json:"sample,omitempty"`

	// 修改时间
	UpdateDate *string `json:"update_date,omitempty"`

	// 采样行数
	RowCount float32 `json:"row_count,omitempty"`

	// 列数
	ColumnCount float32 `json:"column_count,omitempty"`

	// 是否唯一
	Unique *bool `json:"unique,omitempty"`

	// 自动停止
	AutoStop *bool `json:"auto_stop,omitempty"`

	// 时间档案
	TimeProfile *bool `json:"time_profile,omitempty"`

	// duilie
	Queue *string `json:"queue,omitempty"`

	// 连接id
	DwId *string `json:"dw_id,omitempty"`

	// 列概要信息
	ColunmsMetric *interface{} `json:"colunms_metric,omitempty"`

	// 列信息
	ColumnsList *[]ColumnInfo `json:"columns_list,omitempty"`
}

func (o ProfileInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProfileInfo struct{}"
	}

	return strings.Join([]string{"ProfileInfo", string(data)}, " ")
}
