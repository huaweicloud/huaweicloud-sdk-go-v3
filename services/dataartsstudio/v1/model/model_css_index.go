package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CssIndex struct {

	// 索引名称
	IndexName *string `json:"index_name,omitempty"`

	// 索引的guid
	IndexGuid *string `json:"index_guid,omitempty"`

	// 索引的唯一标识名称
	IndexQualifiedName *string `json:"index_qualified_name,omitempty"`

	// 索引中文档总数
	IndexDocCount *int32 `json:"index_doc_count,omitempty"`

	// 索引数据量大小
	IndexDataSize *float64 `json:"index_data_size,omitempty"`
}

func (o CssIndex) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CssIndex struct{}"
	}

	return strings.Join([]string{"CssIndex", string(data)}, " ")
}
