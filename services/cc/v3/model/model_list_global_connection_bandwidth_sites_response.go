package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalConnectionBandwidthSitesResponse Response Object
type ListGlobalConnectionBandwidthSitesResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 站点信息列表响应体。
	SiteInfos      []GlobalConnectionBandwidthSites `json:"site_infos"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListGlobalConnectionBandwidthSitesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalConnectionBandwidthSitesResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalConnectionBandwidthSitesResponse", string(data)}, " ")
}
