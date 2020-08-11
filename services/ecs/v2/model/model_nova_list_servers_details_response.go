/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Response Object
type NovaListServersDetailsResponse struct {
	// 查询云服务器信息列表。
	Servers []NovaServer `json:"servers,omitempty"`
	// 分页查询时，查询下一页数据链接。
	ServersLinks []PageLink `json:"servers_links,omitempty"`
}
