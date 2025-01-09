package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWksEdgeSitesResponse Response Object
type ListWksEdgeSitesResponse struct {

	// 边缘小站列表。
	WksEdgeSites *[]WksEdgeSiteDetail `json:"wks_edge_sites,omitempty"`

	// 站点总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListWksEdgeSitesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWksEdgeSitesResponse struct{}"
	}

	return strings.Join([]string{"ListWksEdgeSitesResponse", string(data)}, " ")
}
