/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

//  
type NovaLink struct {
	// 相应资源的链接。
	Href string `json:"href"`
	// 有三种取值。self：自助链接包含版本链接的资源。立即链接后使用这些链接。bookmark：书签链接提供了一个永久资源的永久链接，该链接适合于长期存储。alternate：备用链接可以包含资源的替换表示形式。例如，OpenStack计算映像可能在OpenStack映像服务中有一个替代表示。
	Rel string `json:"rel"`
}
