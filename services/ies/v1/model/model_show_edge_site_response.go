package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEdgeSiteResponse Response Object
type ShowEdgeSiteResponse struct {
	EdgeSite       *EdgeSiteDetail `json:"edge_site,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowEdgeSiteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeSiteResponse struct{}"
	}

	return strings.Join([]string{"ShowEdgeSiteResponse", string(data)}, " ")
}
