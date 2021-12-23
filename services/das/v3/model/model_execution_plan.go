package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecutionPlan struct {
	// id

	Id string `json:"id"`
	// select子句的类型

	SelectType string `json:"select_type"`
	// 数据库表

	Table string `json:"table"`
	// 查询将匹配记录的分区

	Partitions string `json:"partitions"`
	// 访问类型

	Type string `json:"type"`
	// 可能使用的键(索引)

	PossibleKeys string `json:"possible_keys"`
	// 实际使用的键(索引)

	Key string `json:"key"`
	// 决定使用的键的长度

	KeyLen string `json:"key_len"`
	// 使用哪个列或常数与键一起来选择行

	Ref string `json:"ref"`
	// MySQL认为它执行查询时必须检查的行数

	Rows string `json:"rows"`
	// 按表条件过滤的表行的估计百分比

	Filtered string `json:"filtered"`
	// 其他信息

	Extra string `json:"extra"`
}

func (o ExecutionPlan) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionPlan struct{}"
	}

	return strings.Join([]string{"ExecutionPlan", string(data)}, " ")
}
