package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FullSqlTask struct {

	// 任务ID。
	Id int64 `json:"id"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 实例名称。
	InstanceName string `json:"instance_name"`

	// 节点ID。
	NodeId string `json:"node_id"`

	// 开始时间（Unix timestamp），单位：毫秒。
	StartAt int64 `json:"start_at"`

	// 结束时间（Unix timestamp），单位：毫秒。
	EndAt int64 `json:"end_at"`

	// 批次ID。
	BatchId string `json:"batch_id"`

	// 用户名。
	UserList []string `json:"user_list"`

	// 关键字。
	Keyword []string `json:"keyword"`

	// 数据库。
	DbList *[]string `json:"db_list,omitempty"`

	// 操作类型。
	OperationList []string `json:"operation_list"`

	// 线程ID。
	ThreadIdList []string `json:"thread_id_list"`

	// 事务ID。
	TrxIdList []string `json:"trx_id_list"`

	// 执行状态（0:成功，1:失败）。
	StatusList []string `json:"status_list"`

	// SQL模板ID。
	SqlTemplateIds []string `json:"sql_template_ids"`

	// 任务状态（0:等待解析，1:解析中，2:解析完成，-1:失败）。
	Status int32 `json:"status"`

	// 任务进度，取值为0-100。
	Progress float64 `json:"progress"`

	// 失败原因。
	Reason string `json:"reason"`

	// 创建时间（Unix timestamp），单位：毫秒。
	CreateAt int64 `json:"create_at"`

	// 更新时间（Unix timestamp），单位：毫秒。
	UpdateAt int64 `json:"update_at"`
}

func (o FullSqlTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullSqlTask struct{}"
	}

	return strings.Join([]string{"FullSqlTask", string(data)}, " ")
}
