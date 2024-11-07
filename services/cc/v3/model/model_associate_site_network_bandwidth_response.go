package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateSiteNetworkBandwidthResponse Response Object
type AssociateSiteNetworkBandwidthResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	SiteConnection *SiteConnection `json:"site_connection"`
	HttpStatusCode int             `json:"-"`
}

func (o AssociateSiteNetworkBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateSiteNetworkBandwidthResponse struct{}"
	}

	return strings.Join([]string{"AssociateSiteNetworkBandwidthResponse", string(data)}, " ")
}
