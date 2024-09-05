package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudServersResponse Response Object
type ListCloudServersResponse struct {

	// 查询云服务器信息列表。
	Servers *[]CloudServer `json:"servers,omitempty"`

	// 分页查询时，查询下一页数据链接。
	ServersLinks *[]PageLink `json:"servers_links,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListCloudServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudServersResponse struct{}"
	}

	return strings.Join([]string{"ListCloudServersResponse", string(data)}, " ")
}
