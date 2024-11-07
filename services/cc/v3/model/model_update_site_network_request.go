package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSiteNetworkRequest Request Object
type UpdateSiteNetworkRequest struct {

	// 分支网络的ID。
	SiteNetworkId string `json:"site_network_id"`

	Body *UpdateSiteNetworkRequestBody `json:"body,omitempty"`
}

func (o UpdateSiteNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSiteNetworkRequest struct{}"
	}

	return strings.Join([]string{"UpdateSiteNetworkRequest", string(data)}, " ")
}
