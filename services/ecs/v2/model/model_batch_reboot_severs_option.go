/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 
type BatchRebootSeversOption struct {
	// 云服务器ID列表。
	Servers []ServerId `json:"servers"`
	// 重启类型：  - SOFT：普通重启。 - HARD：强制重启。
	Type string `json:"type"`
}
