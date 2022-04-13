package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据同步配置库规则信息体
type ConfigTransformationVo struct {
	// 库名.表名。

	DbTableName string `json:"db_table_name"`
	// 库名。长度限制256。

	DbName string `json:"db_name"`
	// 表名。长度限制256。

	TableName string `json:"table_name"`
	// 列名。长度限制256。

	ColNames string `json:"col_names"`
	// 主键或唯一索引。长度限制256。

	PrimKeyOrIndex string `json:"prim_key_or_index"`
	// 优化查询所需的索引。长度限制256。

	Indexs string `json:"indexs"`
	// 过滤条件。长度限制256。

	Values string `json:"values"`
}

func (o ConfigTransformationVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigTransformationVo struct{}"
	}

	return strings.Join([]string{"ConfigTransformationVo", string(data)}, " ")
}
