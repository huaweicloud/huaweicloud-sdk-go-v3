package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEdgeSitesResponse Response Object
type ListEdgeSitesResponse struct {

	// 边缘小站列表。
	EdgeSites *[]EdgeSiteDetail `json:"edge_sites,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListEdgeSitesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEdgeSitesResponse struct{}"
	}

	return strings.Join([]string{"ListEdgeSitesResponse", string(data)}, " ")
}
