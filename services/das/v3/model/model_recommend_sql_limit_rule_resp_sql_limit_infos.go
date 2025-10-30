package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecommendSqlLimitRuleRespSqlLimitInfos struct {
	RawSql *RecommendSqlLimitRuleRespRawSql `json:"raw_sql,omitempty"`

	// 平均时间
	AverageTime *float64 `json:"average_time,omitempty"`

	// 数量
	Count *float64 `json:"count,omitempty"`

	// mysql 提供， taurus不提供
	MaxTime *int64 `json:"maxTime,omitempty"`

	// 执行时间
	ExeTime *int64 `json:"exe_time,omitempty"`
}

func (o RecommendSqlLimitRuleRespSqlLimitInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecommendSqlLimitRuleRespSqlLimitInfos struct{}"
	}

	return strings.Join([]string{"RecommendSqlLimitRuleRespSqlLimitInfos", string(data)}, " ")
}
