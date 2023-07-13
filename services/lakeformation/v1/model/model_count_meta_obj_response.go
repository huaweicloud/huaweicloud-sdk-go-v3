package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountMetaObjResponse Response Object
type CountMetaObjResponse struct {

	// 总数量
	TotalCount *int64 `json:"total_count,omitempty"`

	// 函数数量
	FunctionsCount *int64 `json:"functions_count,omitempty"`

	// 目录数量
	CatalogsCount *int64 `json:"catalogs_count,omitempty"`

	// 数据库数量
	DatabasesCount *int64 `json:"databases_count,omitempty"`

	// 表数量
	TablesCount *int64 `json:"tables_count,omitempty"`

	// 分区数量
	PartitionsCount *int64 `json:"partitions_count,omitempty"`

	// 索引数量
	IndexesCount   *int64 `json:"indexes_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountMetaObjResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountMetaObjResponse struct{}"
	}

	return strings.Join([]string{"CountMetaObjResponse", string(data)}, " ")
}
