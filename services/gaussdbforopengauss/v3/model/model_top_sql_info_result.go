package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopSqlInfoResult Top SQL信息
type TopSqlInfoResult struct {

	// **参数解释**: Top SQL的归一化SQL ID。 **取值范围**: 不涉及。
	SqlId *string `json:"sql_id,omitempty"`

	// **参数解释**: Top SQL的用户名。 **取值范围**: 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: Top SQL的SQL文本。 **取值范围**: 不涉及。
	SqlText *string `json:"sql_text,omitempty"`

	// **参数解释**: Top SQL的调用频率占比（%）。 **取值范围**: 0-100。
	CallsPercent *string `json:"calls_percent,omitempty"`

	// **参数解释**: Top SQL的调用频率占比。 **取值范围**: 0-100。
	CpuPercent *string `json:"cpu_percent,omitempty"`

	// **参数解释**: Top SQL的IO开销占比。 **取值范围**: 0-100。
	IoPercent *string `json:"io_percent,omitempty"`

	// **参数解释** Top SQL的调用次数。 **取值范围** 大于等于0。
	Calls *string `json:"calls,omitempty"`

	// **参数解释** Top SQL的返回元组数。 **取值范围** 大于等于0。
	ReturnedRows *string `json:"returned_rows,omitempty"`

	// **参数解释** Top SQL的读取元组数。 **取值范围** 大于等于0。
	TupleRead *string `json:"tuple_read,omitempty"`

	// **参数解释** Top SQL的平均时间开销。单位ms。 **取值范围** 大于等于0。
	AvgElapseTime *string `json:"avg_elapse_time,omitempty"`

	// **参数解释** Top SQL的总时间开销。单位ms。 **取值范围** 大于等于0。
	TotalElapseTime *string `json:"total_elapse_time,omitempty"`

	// **参数解释**: Top SQL的CPU开销。单位ms。 **取值范围**: 不涉及。
	CpuTime *string `json:"cpu_time,omitempty"`

	// **参数解释**: Top SQL的IO开销。单位ms。 **取值范围**: 不涉及。
	IoTime *string `json:"io_time,omitempty"`

	// **参数解释**: Top SQL的最小执行时间。单位ms。 **取值范围**: 不涉及。
	MinElapseTime *string `json:"min_elapse_time,omitempty"`

	// **参数解释**: Top SQL最大执行时间。单位ms。 **取值范围**: 不涉及。
	MaxElapseTime *string `json:"max_elapse_time,omitempty"`

	// **参数解释** Top SQL的SQL命中率。 **取值范围** 大于等于0。
	SqlHitRatio *string `json:"sql_hit_ratio,omitempty"`

	// **参数解释**: Top SQL的节点ID。 **取值范围**: 不涉及。
	NodeId *string `json:"node_id,omitempty"`

	// **参数解释**: Top SQL的数据库名（引擎版本8.200及以上支持）。 **取值范围**: 不涉及。
	DbName *string `json:"db_name,omitempty"`

	// **参数解释**: Top SQL的节点名称。 **取值范围**: 不涉及。
	NodeName *string `json:"node_name,omitempty"`
}

func (o TopSqlInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopSqlInfoResult struct{}"
	}

	return strings.Join([]string{"TopSqlInfoResult", string(data)}, " ")
}
