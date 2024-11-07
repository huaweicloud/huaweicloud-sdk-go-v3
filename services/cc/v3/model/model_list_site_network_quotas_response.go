package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSiteNetworkQuotasResponse Response Object
type ListSiteNetworkQuotasResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	// 分支网络的配额列表。
	Quotas         []SiteNetworkQuota `json:"quotas"`
	HttpStatusCode int                `json:"-"`
}

func (o ListSiteNetworkQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSiteNetworkQuotasResponse struct{}"
	}

	return strings.Join([]string{"ListSiteNetworkQuotasResponse", string(data)}, " ")
}
