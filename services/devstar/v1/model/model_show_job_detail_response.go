/*
    * Devstar
    *
    * Devstar API
    *
*/

package model

// Response Object
type ShowJobDetailResponse struct {
	// 任务的id
	Id string `json:"id,omitempty"`
	// 任务的名称
	Name string `json:"name,omitempty"`
	// 任务的状态
	JobStatus map[string]interface{} `json:"job_status,omitempty"`
	// 任务结果信息
	JobResult string `json:"job_result,omitempty"`
}
