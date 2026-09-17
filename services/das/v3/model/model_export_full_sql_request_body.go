package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportFullSqlRequestBody Console Full Sql Export请求体
type ExportFullSqlRequestBody struct {

	// 开始时间（Unix timestamp），单位：毫秒
	StartAt int64 `json:"start_at"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndAt int64 `json:"end_at"`

	// SQL洞察任务ID列表
	TaskIds *[]int64 `json:"task_ids,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 关键字
	Keyword *string `json:"keyword,omitempty"`

	// 是否模糊匹配
	Fuzzy *bool `json:"fuzzy,omitempty"`

	// 用户名
	UserList *[]string `json:"user_list,omitempty"`

	// 数据库
	DbList *[]string `json:"db_list,omitempty"`

	// 操作类型
	OperationList *[]string `json:"operation_list,omitempty"`

	// 客户端IP
	ClientIpList *[]string `json:"client_ip_list,omitempty"`

	// 线程ID
	ThreadIdList *[]int64 `json:"thread_id_list,omitempty"`

	// 事务ID
	TrxIdList *[]int64 `json:"trx_id_list,omitempty"`

	// 会话ID
	SessionIdList *[]int64 `json:"session_id_list,omitempty"`

	// 执行状态（0：成功，1：失败）
	StatusList *[]int32 `json:"status_list,omitempty"`

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

	// OBS桶名
	BucketName *string `json:"bucket_name,omitempty"`

	// 导出的列名
	ExportColumnList *[]string `json:"export_column_list,omitempty"`

	// 时区
	TimeZone *string `json:"time_zone,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 任务ID
	TaskId *int64 `json:"task_id,omitempty"`
}

func (o ExportFullSqlRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFullSqlRequestBody struct{}"
	}

	return strings.Join([]string{"ExportFullSqlRequestBody", string(data)}, " ")
}
