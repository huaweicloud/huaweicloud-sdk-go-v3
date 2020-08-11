/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 
type BatchStopServersOption struct {
	// 标记为启动云服务器操作。
	Servers []ServerId `json:"servers"`
	// 关机类型，默认为SOFT：  - SOFT：普通关机（默认）。 - HARD：强制关机。
	Type string `json:"type,omitempty"`
}
