/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 
type BatchStartServersOption struct {
	// 云服务器ID列表
	Servers []ServerId `json:"servers"`
}
