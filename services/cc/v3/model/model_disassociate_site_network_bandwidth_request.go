package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateSiteNetworkBandwidthRequest Request Object
type DisassociateSiteNetworkBandwidthRequest struct {

	// 分支网络的ID。
	SiteNetworkId string `json:"site_network_id"`

	// 实例ID。
	SiteConnectionId string `json:"site_connection_id"`
}

func (o DisassociateSiteNetworkBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateSiteNetworkBandwidthRequest struct{}"
	}

	return strings.Join([]string{"DisassociateSiteNetworkBandwidthRequest", string(data)}, " ")
}
