package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSiteNetworkBandwidthResponse Response Object
type UpdateSiteNetworkBandwidthResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	SiteConnection *SiteConnection `json:"site_connection"`
	HttpStatusCode int             `json:"-"`
}

func (o UpdateSiteNetworkBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSiteNetworkBandwidthResponse struct{}"
	}

	return strings.Join([]string{"UpdateSiteNetworkBandwidthResponse", string(data)}, " ")
}
