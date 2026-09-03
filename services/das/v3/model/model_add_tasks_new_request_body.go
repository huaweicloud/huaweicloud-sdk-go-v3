package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddTasksNewRequestBody 创建多天全量SQL明细解析任务请求体
type AddTasksNewRequestBody struct {

	// 实例ID，实例的唯一标识
	InstanceId string `json:"instance_id"`

	// 开始时间（Unix timestamp），单位：毫秒
	StartAt int64 `json:"start_at"`

	// 结束时间（Unix timestamp），单位：毫秒
	EndAt int64 `json:"end_at"`

	// 节点ID，实例节点的唯一标识
	NodeId *string `json:"node_id,omitempty"`

	// 关键字，可组合，用逗号分隔
	Keyword *string `json:"keyword,omitempty"`

	// 用户名，可组合，用逗号分隔
	UserList *string `json:"user_list,omitempty"`

	// 数据库，可组合，用逗号分隔
	DbList *string `json:"db_list,omitempty"`

	// 操作类型，可组合，用逗号分隔
	OperationList *string `json:"operation_list,omitempty"`

	// 线程ID，可组合，用逗号分隔
	ThreadIdList *string `json:"thread_id_list,omitempty"`

	// 事务ID，可组合，用逗号分隔
	TrxIdList *string `json:"trx_id_list,omitempty"`

	// 执行状态，可组合，用逗号分隔
	StatusList *string `json:"status_list,omitempty"`

	// SQL模板ID，可组合，用逗号分隔
	SqlTemplateIds *string `json:"sql_template_ids,omitempty"`
}

func (o AddTasksNewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTasksNewRequestBody struct{}"
	}

	return strings.Join([]string{"AddTasksNewRequestBody", string(data)}, " ")
}
