package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSiteNetworkResponse Response Object
type ShowSiteNetworkResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	SiteNetwork    *SiteNetworkEntry `json:"site_network"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowSiteNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSiteNetworkResponse struct{}"
	}

	return strings.Join([]string{"ShowSiteNetworkResponse", string(data)}, " ")
}
