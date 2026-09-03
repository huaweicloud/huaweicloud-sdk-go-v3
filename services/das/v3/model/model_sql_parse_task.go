package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlParseTask SQL解析任务
type SqlParseTask struct {

	// 任务ID
	Id *int64 `json:"id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名
	InstanceName *string `json:"instance_name,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 开始时间（Unix timestamp），单位：毫秒
	StartAt *int64 `json:"start_at,omitempty"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndAt *int64 `json:"end_at,omitempty"`

	// 批次ID
	BatchId *string `json:"batch_id,omitempty"`

	// 用户名
	UserList *[]string `json:"user_list,omitempty"`

	// 关键字
	Keyword *[]string `json:"keyword,omitempty"`

	// 数据库
	DbList *[]string `json:"db_list,omitempty"`

	// 操作类型
	OperationList *[]string `json:"operation_list,omitempty"`

	// 线程ID
	ThreadIdList *[]string `json:"thread_id_list,omitempty"`

	// 事务ID
	TrxIdList *[]string `json:"trx_id_list,omitempty"`

	// 执行状态
	StatusList *[]string `json:"status_list,omitempty"`

	// SQL模板ID
	SqlTemplateIds *[]string `json:"sql_template_ids,omitempty"`

	// 任务状态
	Status *int32 `json:"status,omitempty"`

	// 创建时间（Unix timestamp），单位：毫秒
	CreateAt *int64 `json:"create_at,omitempty"`

	// 更新时间（Unix timestamp），单位：毫秒
	UpdateAt *int64 `json:"update_at,omitempty"`

	// 任务进度
	Progress *float64 `json:"progress,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`
}

func (o SqlParseTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlParseTask struct{}"
	}

	return strings.Join([]string{"SqlParseTask", string(data)}, " ")
}
