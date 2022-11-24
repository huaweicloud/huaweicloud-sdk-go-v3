package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateEdgeSiteResponse struct {
	EdgeSite       *EdgeSiteDetail `json:"edge_site,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o UpdateEdgeSiteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeSiteResponse struct{}"
	}

	return strings.Join([]string{"UpdateEdgeSiteResponse", string(data)}, " ")
}
