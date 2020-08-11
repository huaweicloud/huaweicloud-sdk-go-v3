/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 弹性云服务器系统标签。
type ServerSystemTag struct {
	// 系统标签的Key值。
	Key string `json:"key,omitempty"`
	// 系统标签的value值。
	Value string `json:"value,omitempty"`
}
