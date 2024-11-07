package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSiteNetworkBandwidthSizeRequest Request Object
type UpdateSiteNetworkBandwidthSizeRequest struct {

	// 分支网络的ID。
	SiteNetworkId string `json:"site_network_id"`

	// 实例ID。
	SiteConnectionId string `json:"site_connection_id"`

	Body *UpdateSiteConnectionBandwidthSizeRequestBody `json:"body,omitempty"`
}

func (o UpdateSiteNetworkBandwidthSizeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSiteNetworkBandwidthSizeRequest struct{}"
	}

	return strings.Join([]string{"UpdateSiteNetworkBandwidthSizeRequest", string(data)}, " ")
}
