package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateSiteNetworkBandwidthResponse Response Object
type DisassociateSiteNetworkBandwidthResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	SiteConnection *SiteConnection `json:"site_connection"`
	HttpStatusCode int             `json:"-"`
}

func (o DisassociateSiteNetworkBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateSiteNetworkBandwidthResponse struct{}"
	}

	return strings.Join([]string{"DisassociateSiteNetworkBandwidthResponse", string(data)}, " ")
}
