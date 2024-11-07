package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSiteNetworkRequest Request Object
type ShowSiteNetworkRequest struct {

	// 分支网络的ID。
	SiteNetworkId string `json:"site_network_id"`
}

func (o ShowSiteNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSiteNetworkRequest struct{}"
	}

	return strings.Join([]string{"ShowSiteNetworkRequest", string(data)}, " ")
}
