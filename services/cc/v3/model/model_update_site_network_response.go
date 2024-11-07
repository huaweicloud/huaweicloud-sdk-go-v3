package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSiteNetworkResponse Response Object
type UpdateSiteNetworkResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	SiteNetwork    *SiteNetworkEntry `json:"site_network"`
	HttpStatusCode int               `json:"-"`
}

func (o UpdateSiteNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSiteNetworkResponse struct{}"
	}

	return strings.Join([]string{"UpdateSiteNetworkResponse", string(data)}, " ")
}
