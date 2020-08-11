/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 弹性云服务器所属安全组列表。
type ServerSecurityGroup struct {
	// 安全组名称或者UUID。
	Name string `json:"name"`
	// 安全组ID。
	Id string `json:"id"`
}
