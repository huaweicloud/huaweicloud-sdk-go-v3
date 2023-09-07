package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DbObjectColumnInfo 列加工对象 约束： 列加工提供列级的查询、映射和过滤能力。 编辑列名后，目标数据库的列名为修改后的名称。 列映射名不能和原列名或已存在的映射名相同。 不支持库名、表名带有换行符的列进行映射。 任务再编辑时，已经同步的表不支持修改列信息。 只有勾选的列才会被同步。 MySQL->MySQL、MySQL->GaussDB(for MySQL)不支持分区表的分区字段进行列映射。 GaussDB分区键不可以被过滤。
type DbObjectColumnInfo struct {

	// 数据库库名称。
	DbName *string `json:"db_name,omitempty"`

	// 数据库schema名称。
	SchemaName *string `json:"schema_name,omitempty"`

	// 数据库表名称。
	TableName *string `json:"table_name,omitempty"`

	// 数据库列信息。
	ColumnInfos *[]QueryColumnInfo `json:"column_infos,omitempty"`

	// 数据库列信息总数，与分页无关，仅作为返回体参数
	TotalCount *int32 `json:"total_count,omitempty"`
}

func (o DbObjectColumnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbObjectColumnInfo struct{}"
	}

	return strings.Join([]string{"DbObjectColumnInfo", string(data)}, " ")
}
