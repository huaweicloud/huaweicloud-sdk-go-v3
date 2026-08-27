package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetSiteAuthConfigRequest Request Object
type ResetSiteAuthConfigRequest struct {

	// 站点ID。
	SiteId string `json:"site_id"`
}

func (o ResetSiteAuthConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetSiteAuthConfigRequest struct{}"
	}

	return strings.Join([]string{"ResetSiteAuthConfigRequest", string(data)}, " ")
}
