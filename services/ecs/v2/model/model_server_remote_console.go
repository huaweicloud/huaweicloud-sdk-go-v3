/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 
type ServerRemoteConsole struct {
	// 远程登录的协议。
	Protocol string `json:"protocol"`
	// 远程登录的类型。
	Type string `json:"type"`
	// 远程登录的url。
	Url string `json:"url"`
}
