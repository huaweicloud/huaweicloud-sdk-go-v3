package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFullSqlTasksRequestBody Fullsql Tasks请求体
type ListFullSqlTasksRequestBody struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 起止时间的查询左区间
	RangeLeft *int64 `json:"range_left,omitempty"`

	// 起止时间的查询右区间
	RangeRight *int64 `json:"range_right,omitempty"`

	// 创建时间的查询左区间
	CreateAtLeft *int64 `json:"create_at_left,omitempty"`

	// 创建时间的查询右区间
	CreateAtRight *int64 `json:"create_at_right,omitempty"`

	// 用户名
	User *string `json:"user,omitempty"`

	// 关键字
	Keyword *string `json:"keyword,omitempty"`

	// 数据库
	DbName *string `json:"db_name,omitempty"`

	// 操作
	Operation *string `json:"operation,omitempty"`

	// 线程ID
	ThreadId *string `json:"thread_id,omitempty"`

	// 事务ID
	TrxId *string `json:"trx_id,omitempty"`

	// 执行状态（0：成功，1：失败）
	Status *string `json:"status,omitempty"`

	// SQL模板ID
	SqlTemplateId *string `json:"sql_template_id,omitempty"`

	// 排序字段（create_at, range_start_at, range_end_at）
	SortField *string `json:"sort_field,omitempty"`

	// 排序规则（true：升序，false：降序）
	Asc *bool `json:"asc,omitempty"`

	// 每页记录数
	PageSize *int32 `json:"page_size,omitempty"`

	// 当前页码
	CurPage *int32 `json:"cur_page,omitempty"`
}

func (o ListFullSqlTasksRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFullSqlTasksRequestBody struct{}"
	}

	return strings.Join([]string{"ListFullSqlTasksRequestBody", string(data)}, " ")
}
