package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSiteNetworkResponse Response Object
type DeleteSiteNetworkResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	SiteNetwork    *SiteNetworkEntry `json:"site_network"`
	HttpStatusCode int               `json:"-"`
}

func (o DeleteSiteNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSiteNetworkResponse struct{}"
	}

	return strings.Join([]string{"DeleteSiteNetworkResponse", string(data)}, " ")
}
