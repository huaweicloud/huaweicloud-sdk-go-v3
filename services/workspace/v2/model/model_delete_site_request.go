package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSiteRequest Request Object
type DeleteSiteRequest struct {

	// 站点ID。
	SiteId string `json:"site_id"`
}

func (o DeleteSiteRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSiteRequest struct{}"
	}

	return strings.Join([]string{"DeleteSiteRequest", string(data)}, " ")
}
