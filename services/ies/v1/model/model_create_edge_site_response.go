package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEdgeSiteResponse Response Object
type CreateEdgeSiteResponse struct {
	EdgeSite       *EdgeSiteDetail `json:"edge_site,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o CreateEdgeSiteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeSiteResponse struct{}"
	}

	return strings.Join([]string{"CreateEdgeSiteResponse", string(data)}, " ")
}
