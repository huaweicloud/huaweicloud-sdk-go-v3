package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSiteNetworksResponse Response Object
type ListSiteNetworksResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 分支网络列表。
	SiteNetworks   []SiteNetworkEntry `json:"site_networks"`
	HttpStatusCode int                `json:"-"`
}

func (o ListSiteNetworksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSiteNetworksResponse struct{}"
	}

	return strings.Join([]string{"ListSiteNetworksResponse", string(data)}, " ")
}
