package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IndexAdviceInfo struct {

	// schema名
	SchemaName *string `json:"schema_name,omitempty"`

	// 表名
	TableName *string `json:"table_name,omitempty"`

	// 索引名
	IndexName *string `json:"index_name,omitempty"`

	// 列
	Columns *[]string `json:"columns,omitempty"`

	// 是否唯一
	Unique *bool `json:"unique,omitempty"`

	// 追踪id
	TrackId *string `json:"track_id,omitempty"`

	// 质量
	Quality *interface{} `json:"quality,omitempty"`

	// ddl需要添加的索引
	DdlAddIndex *string `json:"ddl_add_index,omitempty"`
}

func (o IndexAdviceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndexAdviceInfo struct{}"
	}

	return strings.Join([]string{"IndexAdviceInfo", string(data)}, " ")
}
