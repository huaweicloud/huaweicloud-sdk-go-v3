/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Response Object
type UpdateServerMetadataResponse struct {
	// 用户自定义metadata键值对。
	Metadata map[string]string `json:"metadata,omitempty"`
}
