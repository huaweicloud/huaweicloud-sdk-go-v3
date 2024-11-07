package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSiteNetworkRequestBody 更新分支网络的请求体。
type UpdateSiteNetworkRequestBody struct {
	SiteNetwork *UpdateSiteNetwork `json:"site_network"`
}

func (o UpdateSiteNetworkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSiteNetworkRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSiteNetworkRequestBody", string(data)}, " ")
}
