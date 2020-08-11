/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Response Object
type ShowServerTagsResponse struct {
	// 标签列表
	Tags []ServerTag `json:"tags,omitempty"`
}
