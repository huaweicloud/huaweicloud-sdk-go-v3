package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryUserSelectedObjectInfoReq 查询同步映射请求参数。
type QueryUserSelectedObjectInfoReq struct {

	// 库名。
	DbName *string `json:"db_name,omitempty"`

	// 模式名。
	SchemaName *string `json:"schema_name,omitempty"`

	// 表名。
	TableName *string `json:"table_name,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset 大于等于 0。默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。默认为10，取值范围【1-1000】
	Limit *int32 `json:"limit,omitempty"`

	// 是否有映射。
	HasColumnInfo *bool `json:"has_column_info,omitempty"`
}

func (o QueryUserSelectedObjectInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryUserSelectedObjectInfoReq struct{}"
	}

	return strings.Join([]string{"QueryUserSelectedObjectInfoReq", string(data)}, " ")
}
