/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 云服务器规格。
type SimpleFlavor struct {
	// 云服务器规格的ID。
	Id string `json:"id"`
	// 规格相关快捷链接地址。
	Links []Link `json:"links"`
}
