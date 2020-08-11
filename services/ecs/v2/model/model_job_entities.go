/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 
type JobEntities struct {
	// 每个子任务的执行信息。
	SubJobs []SubJob `json:"sub_jobs,omitempty"`
	// 子任务数量。
	SubJobsTotal int32 `json:"sub_jobs_total,omitempty"`
}
