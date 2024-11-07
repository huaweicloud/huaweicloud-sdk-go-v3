package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSiteNetworkBandwidthSizeResponse Response Object
type UpdateSiteNetworkBandwidthSizeResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	SiteConnection *SiteConnection `json:"site_connection"`
	HttpStatusCode int             `json:"-"`
}

func (o UpdateSiteNetworkBandwidthSizeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSiteNetworkBandwidthSizeResponse struct{}"
	}

	return strings.Join([]string{"UpdateSiteNetworkBandwidthSizeResponse", string(data)}, " ")
}
