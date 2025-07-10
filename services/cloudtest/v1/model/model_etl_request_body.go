package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EtlRequestBody struct {

	// 分页偏移量
	Offset int32 `json:"offset"`

	// 分页大小，最大值1000
	Limit int32 `json:"limit"`

	// 需要同步的表名称
	TableName string `json:"table_name"`

	// 是否是bak表数据
	IsBak *string `json:"is_bak,omitempty"`

	// 查询时间段起始时间
	StartTime *string `json:"start_time,omitempty"`

	// 查询时间段截止时间
	EndTime *string `json:"end_time,omitempty"`

	// 用于时间段过滤的字段
	FilterTimeField *string `json:"filter_time_field,omitempty"`

	// 用于时间排序的字段，不传默认按照更新时间排序
	SortField *string `json:"sort_field,omitempty"`

	// 分库编号（数字类型）
	SchemaNo string `json:"schema_no"`
}

func (o EtlRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EtlRequestBody struct{}"
	}

	return strings.Join([]string{"EtlRequestBody", string(data)}, " ")
}
