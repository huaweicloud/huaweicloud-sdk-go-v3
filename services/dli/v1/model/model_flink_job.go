package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkJob 作业列表信息。
type FlinkJob struct {

	// 作业ID。
	JobId int64 `json:"job_id"`

	// 作业名称。长度限制：0-57个字符。
	Name *string `json:"name,omitempty"`

	// 作业描述。长度限制：0-2048个字符。
	Desc *string `json:"desc,omitempty"`

	// 用户名，当“show_detail”为“false”时独有。
	UserName *string `json:"user_name,omitempty"`

	// 作业类型。
	JobType *string `json:"job_type,omitempty"`

	// 作业状态。
	Status *string `json:"status,omitempty"`

	// 作业状态描述。
	StatusDesc *string `json:"status_desc,omitempty"`

	// 作业创建时间。
	CreateTime int64 `json:"create_time"`

	// 作业开始时间。
	StartTime *int64 `json:"start_time,omitempty"`

	// 作业运行时长, 单位ms，当“show_detail”为“false”时独有。
	Duration *int64 `json:"duration,omitempty"`

	// 父作业ID，“show_detail”为“false”时独有。
	RootId *int64 `json:"root_id,omitempty"`

	// 作业所属用户标识，“show_detail”为“true”时独有。
	UserId *string `json:"user_id,omitempty"`

	// 作业所属项目标识，“show_detail”为“true”时独有。
	ProjectId *string `json:"project_id,omitempty"`

	// Stream SQL语句，“show_detail”为“false”时独有。
	SqlBody *string `json:"sql_body,omitempty"`

	// 作业运行模式： shared_cluster：共享。 exclusive_cluster：独享。 edge_node：边缘节点。 show_detail为true时独有.
	RunMode *string `json:"run_mode,omitempty"`

	// jar包主类，“show_detail”为“false”时独有。
	MainClass *string `json:"main_class,omitempty"`

	// jar包作业运行参数，多个参数之间用空格分隔。show_detail为true时独有的。
	EntrypointArgs *string `json:"entrypoint_args,omitempty"`

	// 作业执行计划，“show_detail”为“false”时独有。
	ExecutionGraph *string `json:"execution_graph,omitempty"`

	// 作业更新时间，“show_detail”为“false”时独有。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 作业的流图是否可编辑。“true”表示作业的流图可以编辑，“false”表示作业的流图不可以编辑。
	GraphEditorEnabled *bool `json:"graph_editor_enabled,omitempty"`

	// 作业是否有保存点。“true”表示作业有保存点，“false”表示作业没有保存点。
	HasSavepoint *bool `json:"has_savepoint,omitempty"`

	// 队列名字。
	QueueName *string `json:"queue_name,omitempty"`

	// 边缘计算组ID列表。
	EdgeGroupIds *[]string `json:"edge_group_ids,omitempty"`

	// 重启次数。
	RestartTimes *int64 `json:"restart_times,omitempty"`

	// 保存点路径。
	SavepointPath *string `json:"savepoint_path,omitempty"`

	JobConfig *FlinkJobConfig `json:"job_config,omitempty"`
}

func (o FlinkJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJob struct{}"
	}

	return strings.Join([]string{"FlinkJob", string(data)}, " ")
}
