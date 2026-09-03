package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchNewRequest Request Object
type SearchNewRequest struct {

	// 实例ID，实例的唯一标识
	InstanceId string `json:"instance_id"`

	// 开始时间（Unix timestamp），单位：毫秒
	StartAt int64 `json:"start_at"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndAt int64 `json:"end_at"`

	// SQL洞察任务ID，与task_ids参数二选一
	TaskId *int64 `json:"task_id,omitempty"`

	// SQL洞察任务ID列表，时间范围大于1天的SQL洞察任务后台将拆分为多个任务解析，与task_id参数二选一
	TaskIds *[]int64 `json:"task_ids,omitempty"`

	// 节点ID，实例节点的唯一标识
	NodeId *string `json:"node_id,omitempty"`

	// 关键字，可组合，用逗号分隔
	Keyword *string `json:"keyword,omitempty"`

	// 是否模糊匹配
	Fuzzy *string `json:"fuzzy,omitempty"`

	// 用户名，可组合，用逗号分隔
	UserList *string `json:"user_list,omitempty"`

	// 数据库，可组合，用逗号分隔
	DbList *string `json:"db_list,omitempty"`

	// 操作类型，可组合，用逗号分隔
	OperationList *string `json:"operation_list,omitempty"`

	// 客户端IP，可组合，用逗号分隔
	ClientIpList *string `json:"client_ip_list,omitempty"`

	// 线程ID，可组合，用逗号分隔
	ThreadIdList *string `json:"thread_id_list,omitempty"`

	// 事务ID，可组合，用逗号分隔
	TrxIdList *string `json:"trx_id_list,omitempty"`

	// 会话ID，可组合，用逗号分隔
	SessionIdList *string `json:"session_id_list,omitempty"`

	// 执行状态，可组合，用逗号分隔
	StatusList *string `json:"status_list,omitempty"`

	// SQL模板ID，可组合，用逗号分隔
	SqlTemplateIds *string `json:"sql_template_ids,omitempty"`

	// 最小执行耗时（毫秒）
	CostMin *float64 `json:"cost_min,omitempty"`

	// 最大执行耗时（毫秒）
	CostMax *float64 `json:"cost_max,omitempty"`

	// 最小扫描行数
	ScanMin *int64 `json:"scan_min,omitempty"`

	// 最大扫描行数
	ScanMax *int64 `json:"scan_max,omitempty"`

	// 最小影响行数
	AffectMin *int64 `json:"affect_min,omitempty"`

	// 最大影响行数
	AffectMax *int64 `json:"affect_max,omitempty"`

	// 最小返回行数
	ReturnMin *int64 `json:"return_min,omitempty"`

	// 最大返回行数
	ReturnMax *int64 `json:"return_max,omitempty"`

	// 排序字段，取值范围：execute_at（执行时间）、execute_cost（执行耗时）、lock_wait_time（锁等待时间）、rows_examined（扫描行数）、rows_returned（返回行数）
	SortField *string `json:"sort_field,omitempty"`

	// 排序顺序，true（正序）、false（逆序）
	Asc *bool `json:"asc,omitempty"`

	// 页码
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页记录数
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o SearchNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchNewRequest struct{}"
	}

	return strings.Join([]string{"SearchNewRequest", string(data)}, " ")
}
