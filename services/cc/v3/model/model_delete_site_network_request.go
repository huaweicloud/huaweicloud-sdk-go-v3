package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSiteNetworkRequest Request Object
type DeleteSiteNetworkRequest struct {

	// 分支网络的ID。
	SiteNetworkId string `json:"site_network_id"`
}

func (o DeleteSiteNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSiteNetworkRequest struct{}"
	}

	return strings.Join([]string{"DeleteSiteNetworkRequest", string(data)}, " ")
}
