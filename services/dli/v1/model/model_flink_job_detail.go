package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlinkJobDetail 作业详细信息。
type FlinkJobDetail struct {

	// 作业ID。
	JobId int64 `json:"job_id"`

	// 作业名称。长度限制：0-57个字符。
	Name *string `json:"name,omitempty"`

	// 作业描述。长度限制：0-2048个字符。
	Desc *string `json:"desc,omitempty"`

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

	// 作业所属用户标识。
	UserId *string `json:"user_id,omitempty"`

	// 队列名称。长度限制：1-128个字符。
	QueueName *string `json:"queue_name,omitempty"`

	// 作业所属项目标识。
	ProjectId *string `json:"project_id,omitempty"`

	// Stream SQL语句。
	SqlBody *string `json:"sql_body,omitempty"`

	// 作业运行模式： shared_cluster：共享。 exclusive_cluster：独享。 edge_node：边缘节点。
	RunMode *string `json:"run_mode,omitempty"`

	JobConfig *FlinkJobConfig `json:"job_config,omitempty"`

	// jar包主类。
	MainClass *string `json:"main_class,omitempty"`

	// jar包作业运行参数，多个参数之间空格分隔。
	EntrypointArgs *string `json:"entrypoint_args,omitempty"`

	// 作业执行计划。
	ExecutionGraph *string `json:"execution_graph,omitempty"`

	// 作业更新时间。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 手动产生的Checkpoint的保存路径。
	SavepointPath *string `json:"savepoint_path,omitempty"`
}

func (o FlinkJobDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlinkJobDetail struct{}"
	}

	return strings.Join([]string{"FlinkJobDetail", string(data)}, " ")
}
