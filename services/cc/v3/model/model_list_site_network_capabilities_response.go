package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSiteNetworkCapabilitiesResponse Response Object
type ListSiteNetworkCapabilitiesResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 分支网络租户能力列表。
	Capabilities   []SiteNetworkCapabilityEntry `json:"capabilities"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListSiteNetworkCapabilitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSiteNetworkCapabilitiesResponse struct{}"
	}

	return strings.Join([]string{"ListSiteNetworkCapabilitiesResponse", string(data)}, " ")
}
