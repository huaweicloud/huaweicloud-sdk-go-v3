package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObjectsHealthCompareOverviewInfo 对比结果信息体。
type ObjectsHealthCompareOverviewInfo struct {

	// 对象类型。取值： - DB：数据库。 - TABLE：表。 - VIEW：视图。 - EVENT：事件。 - ROUTINE：存储过程和函数。 - INDEX：索引。 - TRIGGER：触发器。 - SYNONYM：同义词。 - FUNCTION：函数。 - PROCEDURE：存储过程。 - TYPE：自定义类型。 - RULE：规则。 - DEFAULT_TYPE：缺省值。 - PLAN_GUIDE：执行计划。 - CONSTRAINT：约束。 - FILE_GROUP：文件组。 - PARTITION_FUNCTION：分区函数。 - PARTITION_SCHEME：分区方案。 - TABLE_COLLATION：表的排序规则。 - EXTENSIONS：插件。
	Type *string `json:"type,omitempty"`

	// 源数量。
	SourceCount *int64 `json:"source_count,omitempty"`

	// 目标数量。
	TargetCount *int64 `json:"target_count,omitempty"`

	// 对比结果： - 0：不一致。 - 2：一致。 - 3：未完成。
	Status *int32 `json:"status,omitempty"`
}

func (o ObjectsHealthCompareOverviewInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectsHealthCompareOverviewInfo struct{}"
	}

	return strings.Join([]string{"ObjectsHealthCompareOverviewInfo", string(data)}, " ")
}
