package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObjectsCompareOverviewInfo 对比结果信息体。
type ObjectsCompareOverviewInfo struct {

	// 对象类型： 取值： - DB：数据库。 - TABLE：表。 - VIEW：视图。 - EVENT：事件。 - ROUTINE：存储过程和函数。 - INDEX：索引。 - TRIGGER：触发器。 - SYNONYM：同义词。 - FUNCTION：函数。 - PROCEDURE：存储过程。 - TYPE：自定义类型。 - RULE：规则。 - DEFAULT_TYPE：缺省值。 - PLAN_GUIDE：执行计划。 - CONSTRAINT：约束。 - FILE_GROUP：文件组。 - PARTITION_FUNCTION：分区函数。 - PARTITION_SCHEME：分区方案。 - TABLE_COLLATION：表的排序规则。 - EXTENSIONS：插件。
	Type *string `json:"type,omitempty"`

	// 该类型的对象在源库的个数。
	SourceCount *int64 `json:"source_count,omitempty"`

	// 该类型的对象在目标库的个数。
	TargetCount *int64 `json:"target_count,omitempty"`

	// 对比结果，0为不一致，2为一致，3为未完成。
	Status *int32 `json:"status,omitempty"`
}

func (o ObjectsCompareOverviewInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectsCompareOverviewInfo struct{}"
	}

	return strings.Join([]string{"ObjectsCompareOverviewInfo", string(data)}, " ")
}
