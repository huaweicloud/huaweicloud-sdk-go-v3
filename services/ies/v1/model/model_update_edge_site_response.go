package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEdgeSiteResponse Response Object
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
