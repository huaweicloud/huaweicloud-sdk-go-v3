/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// This is a auto create Body Object
type BatchAddServerNicsRequestBody struct {
	// 需要添加的网卡参数列表。
	Nics []BatchAddServerNicOption `json:"nics"`
}
