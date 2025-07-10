package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSiteConfigsResponse Response Object
type ListSiteConfigsResponse struct {

	// 站点信息列表。
	SiteInfos *[]SiteInfo `json:"site_infos,omitempty"`

	// 站点总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSiteConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSiteConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListSiteConfigsResponse", string(data)}, " ")
}
