/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Response Object
type BatchRebootServersResponse struct {
	// 提交任务成功后返回的任务ID，用户可以使用该ID对任务执行情况进行查询。
	JobId string `json:"job_id,omitempty"`
}
