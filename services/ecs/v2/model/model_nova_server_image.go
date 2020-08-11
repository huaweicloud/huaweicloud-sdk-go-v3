/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

//  
type NovaServerImage struct {
	// 镜像ID。
	Id string `json:"id"`
	// 云服务器类型相关标记快捷链接信息。
	Links []NovaLink `json:"links"`
}
