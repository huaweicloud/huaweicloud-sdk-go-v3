/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

//  
type PageLink struct {
	// 相应资源的链接。
	Href string `json:"href"`
	// 对应快捷链接。
	Rel string `json:"rel"`
}
