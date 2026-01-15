package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrationObjectOverviewInfo 迁移详情概览。
type MigrationObjectOverviewInfo struct {

	// 类型。 DATABASE：数据库 SCHEMA：SCHEMA PACKAGE：package TABLE：数据表 COLUMN：列 VIEW：视图 FUNCTION：函数 PROCEDURE：存储过程 ROUTINE：routine TRIGGER：触发器 INDEX：索引 TABLE_INDEX：普通索引，根据表汇聚 TABLE_RENAME_OR_COPY：表重命名或复制 TABLE_STRUCTURE：表结构 EVENT：事件 SYNONYM：同义词,sqlserver特有 TYPE：自定义类型,sqlserver特有 RULE：规则,sqlserver特有 DEFAULT：缺省值,sqlserver特有 PLAN_GUIDE：执行计划,sqlserver特有 FILE_GROUP：文件组,sqlserver特有 PARTITION_FUNCTION：分区函数,sqlserver特有 SHARD_KEY：mongo特有 VALIDATOR：mongo特有 SEQUENCE：序列 MATVIEW：物化视图 PARTITION_SCHEME：分区方案,sqlserver特有 ACCOUNT：账户 EXTENSION：PG 特有的一些对象:插件 AGGREGATE：PG 特有的一些对象:聚合函数 MATERIALIZED_VIEW：PG 特有的一些对象:物化视图 TEXT_SEARCH_DICTIONARY：PG 特有的一些对象:文本搜索字典 CONVERSION：PG 特有的一些对象:类型转换 DATA_TYPE：PG 特有的一些对象:数据类型 TEXT_SEARCH_CONFIGURATION：PG 特有的一些对象:文本搜索配置 STATISTICS_EXTENSION：PG 特有的一些对象:插件统计 MEMBERSHIP：PG 特有的一些对象:用户成员关系 EVENT_TRIGGER：PG 特有的一些对象:事件触发器 COLLATION：PG 特有的一些对象:排序规则 TEXT_SEARCH_PARSER：PG 特有的一些对象:文本搜索解析器 PRIVILEGES：PG 特有的一些对象:权限 FOREIGN_KEY：PG 特有的一些对象:外键 ROLE：权限
	Type *string `json:"type,omitempty"`

	// 待迁移数量。
	SrcCount *string `json:"src_count,omitempty"`

	// 已迁移数量。
	DstCount *string `json:"dst_count,omitempty"`

	// 状态. NOT_START：未启动，TRANSFERING：迁移中，COMPLETED：已完成，FAILED：失败，TRANSFER_WHEN_END：结束后迁移
	Status *string `json:"status,omitempty"`

	// 开始时间。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间。
	EndTime *string `json:"end_time,omitempty"`
}

func (o MigrationObjectOverviewInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationObjectOverviewInfo struct{}"
	}

	return strings.Join([]string{"MigrationObjectOverviewInfo", string(data)}, " ")
}
