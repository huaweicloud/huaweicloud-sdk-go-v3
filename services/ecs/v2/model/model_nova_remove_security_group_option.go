/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

//  
type NovaRemoveSecurityGroupOption struct {
	// 弹性云服务器移除的安全组名称，会对云服务器中配置的网卡生效。
	Name string `json:"name"`
}
