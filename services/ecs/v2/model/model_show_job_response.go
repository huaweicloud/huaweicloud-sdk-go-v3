/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Response Object
type ShowJobResponse struct {
	// 开始时间。
	BeginTime string `json:"begin_time,omitempty"`
	// 查询Job的API请求出现错误时，返回的错误码。
	Code string `json:"code,omitempty"`
	// 结束时间。
	EndTime string `json:"end_time,omitempty"`
	Entities *JobEntities `json:"entities,omitempty"`
	// Job执行失败时的错误码。  Job执行成功后，该值为null。
	ErrorCode string `json:"error_code,omitempty"`
	// Job执行失败时的错误原因。  Job执行成功后，该值为null。
	FailReason string `json:"fail_reason,omitempty"`
	// 异步请求的任务ID。
	JobId string `json:"job_id,omitempty"`
	// 异步请求的任务类型。
	JobType string `json:"job_type,omitempty"`
	// 查询Job的API请求出现错误时，返回的错误消息。
	Message string `json:"message,omitempty"`
	// Job的状态。  - SUCCESS：成功。  - RUNNING：运行中。  - FAIL：失败。  - INIT：正在初始化。
	Status string `json:"status,omitempty"`
}
