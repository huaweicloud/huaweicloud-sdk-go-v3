package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobExeResult struct {

	// 作业ID。
	Id string `json:"id"`

	// 作业创建时间，十三位时间戳。
	CreateAt int64 `json:"create_at"`

	// 作业更新时间，十三位时间戳。
	UpdateAt int64 `json:"update_at"`

	// 项目编号。获取方法，请参见[获取项目ID](https://support.huaweicloud.com/api-mrs/mrs_02_0011.html)。
	TenantId string `json:"tenant_id"`

	// 作业ID。
	JobId string `json:"job_id"`

	// 作业名称。
	JobName string `json:"job_name"`

	// 作业执行开始时间，十三位时间戳。
	StartTime int64 `json:"start_time"`

	// 作业执行结束时间，十三位时间戳。
	EndTime int64 `json:"end_time"`

	// 作业所属集群ID。
	ClusterId string `json:"cluster_id"`

	// 作业执行组ID
	GroupId string `json:"group_id"`

	// 执行程序jar包或sql文件地址。
	JarPath string `json:"jar_path"`

	// 数据输入地址。
	Input string `json:"input"`

	// 数据输出地址。
	Output *string `json:"output,omitempty"`

	// 作业日志存储地址
	JobLog string `json:"job_log"`

	// 作业类型码。  - 1：MapReduce - 2：Spark - 3：Hive Script - 4：HiveSQL（当前不支持） - 5：DistCp - 6：Spark Script - 7：Spark SQL（该接口当前不支持）
	JobType int32 `json:"job_type"`

	// 导入导出数据。
	FileAction string `json:"file_action"`

	// 程序执行的关键参数，该参数由用户程序内的函数指定，MRS只负责参数的传入。该参数可为空。
	Arguments string `json:"arguments"`

	// HQL脚本语句。
	Hql string `json:"hql"`

	// 作业状态编码：  - 1：Terminated表示已终止的作业状态 - 2：Running表示运行中的作业状态 - 3：Completed表示已完成的作业状态 - 4：Abnormal表示异常的作业状态
	JobState int32 `json:"job_state"`

	// 作业最终状态码。  - 0：未完成 - 1：执行错误，终止执行 - 2：执行完成并且成功 - 3：已取消
	JobFinalStatus int32 `json:"job_final_status"`

	// Hive脚本地址。
	HiveScriptPath string `json:"hive_script_path"`

	// 创建作业的用户ID。
	CreateBy string `json:"create_by"`

	// 当前已完成的步骤数。
	FinishedStep int32 `json:"finished_step"`

	// 作业主ID。
	JobMainId string `json:"job_main_id"`

	// 作业步骤ID。
	JobStepId string `json:"job_step_id"`

	// 延迟时间，十三位时间戳。
	PostponeAt int64 `json:"postpone_at"`

	// 作业步骤名。
	StepName string `json:"step_name"`

	// 步骤数量。
	StepNum int32 `json:"step_num"`

	// 任务数量。
	TaskNum int32 `json:"task_num"`

	// 更新作业的用户ID。
	UpdateBy string `json:"update_by"`

	// 作业执行持续时间，单位：秒。
	SpendTime float32 `json:"spend_time"`

	// 步骤序列号。
	StepSeq int32 `json:"step_seq"`

	// 作业执行进度。
	Progress string `json:"progress"`
}

func (o JobExeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobExeResult struct{}"
	}

	return strings.Join([]string{"JobExeResult", string(data)}, " ")
}
