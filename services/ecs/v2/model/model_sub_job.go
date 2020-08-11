/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 
type SubJob struct {
	// Job的状态。  - SUCCESS：成功。  - RUNNING：运行中。  - FAIL：失败。  - INIT：正在初始化。
	Status string `json:"status,omitempty"`
	Entities *SubJobEntities `json:"entities,omitempty"`
	// 子任务的ID。
	JobId string `json:"job_id,omitempty"`
	// 子任务的类型。
	JobType string `json:"job_type,omitempty"`
	// 开始时间。
	BeginTime string `json:"begin_time,omitempty"`
	// 结束时间。
	EndTime string `json:"end_time,omitempty"`
	// Job执行失败时的错误码。  Job执行成功后，该值为null。
	ErrorCode string `json:"error_code,omitempty"`
	// Job执行失败时的错误原因。  Job执行成功后，该值为null。
	FailReason string `json:"fail_reason,omitempty"`
}
