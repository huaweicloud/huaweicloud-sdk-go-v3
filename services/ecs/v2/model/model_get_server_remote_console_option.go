/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

type GetServerRemoteConsoleOption struct {
	// 远程登录协议，请将protocol配置为“vnc”。
	Protocol string `json:"protocol"`
	// 远程登录的类型，请将type配置为“novnc”。
	Type string `json:"type"`
}
