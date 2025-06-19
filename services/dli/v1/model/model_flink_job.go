package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlinkJob 作业列表信息。
type FlinkJob struct {

	// 参数解释:  作业ID 示例: 50320 约束限制:  无 取值范围: 无 默认取值: 无
	JobId int64 `json:"job_id"`

	// 参数解释:  作业名称 示例: test 约束限制:  长度在[0,57]范围内的字符串 取值范围: 无 默认取值: 无
	Name *string `json:"name,omitempty"`

	// 参数解释:  作业描述 示例: 作业描述 约束限制:  长度在[0,2048]范围内的字符串 取值范围: 无 默认取值: 无
	Desc *string `json:"desc,omitempty"`

	// 参数解释:  用户名称 示例: testuser 约束限制:  长度在[1,128]范围内的字符串 取值范围: 无 默认取值: 无
	UserName *string `json:"user_name,omitempty"`

	// 参数解释:  作业类型 示例: flink_jar_job 约束限制:  无 取值范围: flink_sql_job（flink sql作业） flink_opensource_sql_job（flink opensource sql作业） flink_sql_edge_job（flink sql边缘作业） flink_jar_job（flink自定义作业） 默认取值: 无
	JobType *FlinkJobJobType `json:"job_type,omitempty"`

	// 参数解释:  作业状态 示例: job_running 约束限制:  无 取值范围: job_init（草稿） job_submitting（提交中） job_submit_fail（提交失败） job_running（运行中） job_running_exception（运行异常） job_downloading（下载中） job_idle（空闲） job_canceling（停止中） job_cancel_success（已停止） job_cancel_fail（停止失败） job_savepointing（保存点创建中） job_arrearage_stopped（因欠费被停止） job_arrearage_recovering（欠费作业恢复中） job_finish（已完成） 默认取值: 无
	Status *string `json:"status,omitempty"`

	// 参数解释:  用户名称 示例: 作业状态描述 约束限制:  无 取值范围: 无 默认取值: 无
	StatusDesc *string `json:"status_desc,omitempty"`

	// 参数解释:  作业创建时间 示例: 1516952770835 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	CreateTime int64 `json:"create_time"`

	// 参数解释:  作业开始时间 示例: 1516952710740 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	StartTime *int64 `json:"start_time,omitempty"`

	// 参数解释:  作业运行时长，单位ms，当“show_detail”为“false”时独有 示例: 30000 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	Duration *int64 `json:"duration,omitempty"`

	// 参数解释:  父作业ID，“show_detail”为“false”时独有 示例: -1 约束限制:  无 取值范围: 无 默认取值: 无
	RootId *int64 `json:"root_id,omitempty"`

	// 参数解释:  作业所属用户标识，“show_detail”为“true”时独有 示例: ac4eaa303639409c8ab099d55eb1538e 约束限制:  无 取值范围: 无 默认取值: 无
	UserId *string `json:"user_id,omitempty"`

	// 参数解释:  作业所属用户标识，“show_detail”为“true”时独有 示例: 48cc2c48765f481480c7db940d6409d1 约束限制:  无 取值范围: 无 默认取值: 无
	ProjectId *string `json:"project_id,omitempty"`

	// 参数解释:  Stream SQL语句，“show_detail”为“false”时独有 示例: select * from source_table 约束限制:  无 取值范围: 无 默认取值: 无
	SqlBody *string `json:"sql_body,omitempty"`

	// 参数解释:  作业运行模式，show_detail为true时独有 示例: shared_cluster 约束限制:  无 取值范围: shared_cluster（共享） exclusive_cluster（独享） edge_node（边缘节点） 默认取值: 无
	RunMode *string `json:"run_mode,omitempty"`

	// 参数解释:  jar包主类，“show_detail”为“false”时独有 示例: org.apache.spark.examples.streaming.JavaQueueStream 约束限制:  无 取值范围: 无 默认取值: 无
	MainClass *string `json:"main_class,omitempty"`

	// 参数解释:  jar包作业运行参数，多个参数之间用空格分隔。show_detail为true时独有的 示例: custom.dir=/usr custom.prefix=dli 约束限制:  无 取值范围: 无 默认取值: 无
	EntrypointArgs *string `json:"entrypoint_args,omitempty"`

	// 参数解释:  作业执行计划，“show_detail”为“false”时独有 约束限制:  无 取值范围: 无 默认取值: 无
	ExecutionGraph *string `json:"execution_graph,omitempty"`

	// 参数解释:  作业更新时间，“show_detail”为“false”时独有 示例: 1516952770835 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 参数解释:  作业的流图是否可编辑。“true”表示作业的流图可以编辑，“false”表示作业的流图不可以编辑 示例: false 约束限制:  无 取值范围: true,false 默认取值: 无
	GraphEditorEnabled *bool `json:"graph_editor_enabled,omitempty"`

	// 参数解释:  作业是否有保存点。“true”表示作业有保存点，“false”表示作业没有保存点 示例: false 约束限制:  无 取值范围: true,false 默认取值: 无
	HasSavepoint *bool `json:"has_savepoint,omitempty"`

	// 参数解释:  队列名称 示例: flink_17_queue 约束限制:  无 取值范围: 无 默认取值: 无
	QueueName *string `json:"queue_name,omitempty"`

	// 参数解释:  边缘计算组ID列表 示例: 62de1e1c-066e-48a8-a79d-f461a31b2ee1,2eb00f85-99f2-4144-bcb7-d39ff47f9002 约束限制:  无 取值范围: 无 默认取值: 无
	EdgeGroupIds *[]string `json:"edge_group_ids,omitempty"`

	// 参数解释:  重启次数 示例: 0 约束限制:  无 取值范围: 大于等于0的整数 默认取值: 无
	RestartTimes *int64 `json:"restart_times,omitempty"`

	// 参数解释:  保存点路径 示例: obs://cwk/savepoint/ 约束限制:  无 取值范围: 无 默认取值: 无
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

type FlinkJobJobType struct {
	value string
}

type FlinkJobJobTypeEnum struct {
	FLINK_SQL_JOB            FlinkJobJobType
	FLINK_OPENSOURCE_SQL_JOB FlinkJobJobType
	FLINK_SQL_EDGE_JOB       FlinkJobJobType
	FLINK_JAR_JOB            FlinkJobJobType
}

func GetFlinkJobJobTypeEnum() FlinkJobJobTypeEnum {
	return FlinkJobJobTypeEnum{
		FLINK_SQL_JOB: FlinkJobJobType{
			value: "flink_sql_job",
		},
		FLINK_OPENSOURCE_SQL_JOB: FlinkJobJobType{
			value: "flink_opensource_sql_job",
		},
		FLINK_SQL_EDGE_JOB: FlinkJobJobType{
			value: "flink_sql_edge_job",
		},
		FLINK_JAR_JOB: FlinkJobJobType{
			value: "flink_jar_job",
		},
	}
}

func (c FlinkJobJobType) Value() string {
	return c.value
}

func (c FlinkJobJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlinkJobJobType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
