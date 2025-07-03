package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaListServersResponse Response Object
type NovaListServersResponse struct {

	// 查询云服务器信息列表
	Servers *[]NovaSimpleServer `json:"servers,omitempty"`

	// 分页查询时，查询下一页数据链接。
	ServersLinks   *[]PageLink `json:"servers_links,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o NovaListServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListServersResponse struct{}"
	}

	return strings.Join([]string{"NovaListServersResponse", string(data)}, " ")
}
