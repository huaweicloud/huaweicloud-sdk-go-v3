/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// This is a auto create Body Object
type BatchDeleteServerTagsRequestBody struct {
	// 操作标识（仅支持小写）：delete（删除）。
	Action string `json:"action"`
	// 标签列表。
	Tags []ServerTag `json:"tags"`
}
