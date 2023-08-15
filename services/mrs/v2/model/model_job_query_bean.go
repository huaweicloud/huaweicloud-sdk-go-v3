package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobQueryBean struct {

	// 作业ID。
	JobId *string `json:"job_id,omitempty"`

	// 提交作业的用户名称。
	User *string `json:"user,omitempty"`

	// 作业名称。
	JobName *string `json:"job_name,omitempty"`

	// 作业最终结果。 - FAILED：执行失败的作业 - KILLED：执行中被手动终止的作业。 - UNDEFINED：正在执行的作业。 - SUCCEEDED：执行成功的作业。
	JobResult *string `json:"job_result,omitempty"`

	// 作业执行状态。  - FAILED：失败 - KILLED：已终止 - NEW：已创建 - NEW_SAVING：已创建保存中 - SUBMITTED：已提交 - ACCEPTED：已接受 - RUNNING：运行中 - FINISHED：已完成
	JobState *string `json:"job_state,omitempty"`

	// 作业执行进度。
	JobProgress *float32 `json:"job_progress,omitempty"`

	// 作业类型。 - MapReduce - SparkSubmit：SparkPython类型的作业在查询时作业类型请选择SparkSubmit。 - HiveScript - HiveSql - DistCp，导入、导出数据。 - SparkScript - SparkSql - Flink
	JobType *string `json:"job_type,omitempty"`

	// 作业开始执行时间。单位：毫秒。
	StartedTime *int64 `json:"started_time,omitempty"`

	// 作业提交时间。单位：毫秒。
	SubmittedTime *int64 `json:"submitted_time,omitempty"`

	// 作业完成时间。单位：毫秒。
	FinishedTime *int64 `json:"finished_time,omitempty"`

	// 作业执行时长。单位：毫秒。
	ElapsedTime *int64 `json:"elapsed_time,omitempty"`

	// 运行参数。
	Arguments *string `json:"arguments,omitempty"`

	// 实际作业编号。
	LauncherId *string `json:"launcher_id,omitempty"`

	// 配置参数，用于传-d参数。最多为2048字符，不能包含><|'`&!\\特殊字符，可为空。
	Properties *string `json:"properties,omitempty"`

	// 实际作业编号。
	AppId *string `json:"app_id,omitempty"`

	//  日志链接地址。当前仅SparkSubmit作业支持该参数。 该参数基于集群的EIP访问集群中的YARN WebUI页面，用户如果在VPC界面解绑EIP，MRS服务侧数据会因为未更新导致该参数引用旧EIP导致访问失败，可通过对集群重新进行EIP的绑定来修复该问题。
	TrackingUrl *string `json:"tracking_url,omitempty"`

	// 作业的资源对列类型。
	Queue *string `json:"queue,omitempty"`
}

func (o JobQueryBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobQueryBean struct{}"
	}

	return strings.Join([]string{"JobQueryBean", string(data)}, " ")
}
