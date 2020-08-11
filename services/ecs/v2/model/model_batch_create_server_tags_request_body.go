/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// This is a auto create Body Object
type BatchCreateServerTagsRequestBody struct {
	// 操作标识（仅支持小写）：create（创建）。
	Action string `json:"action"`
	// 标签列表。
	Tags []ServerTag `json:"tags"`
}
