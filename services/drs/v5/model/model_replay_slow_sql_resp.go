package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplaySlowSqlResp 回放慢SQL数据结果对象
type ReplaySlowSqlResp struct {

	// SQL语句类型
	ObjectType *string `json:"object_type,omitempty"`

	// SQL语句
	SlowSql *string `json:"slow_sql,omitempty"`

	// 源库执行耗时
	OldTime *string `json:"old_time,omitempty"`

	// 目标库回放执行耗时
	ReplayTime *string `json:"replay_time,omitempty"`
}

func (o ReplaySlowSqlResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplaySlowSqlResp struct{}"
	}

	return strings.Join([]string{"ReplaySlowSqlResp", string(data)}, " ")
}
