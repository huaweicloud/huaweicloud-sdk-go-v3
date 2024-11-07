package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateSiteNetworkBandwidthRequest Request Object
type AssociateSiteNetworkBandwidthRequest struct {

	// 分支网络的ID。
	SiteNetworkId string `json:"site_network_id"`

	// 实例ID。
	SiteConnectionId string `json:"site_connection_id"`

	Body *AssociateSiteConnectionBandwidthRequestBody `json:"body,omitempty"`
}

func (o AssociateSiteNetworkBandwidthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateSiteNetworkBandwidthRequest struct{}"
	}

	return strings.Join([]string{"AssociateSiteNetworkBandwidthRequest", string(data)}, " ")
}
