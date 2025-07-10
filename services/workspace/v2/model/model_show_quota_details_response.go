package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotaDetailsResponse Response Object
type ShowQuotaDetailsResponse struct {

	// 配额资源列表。
	Resources []ResourceNoLimit `json:"resources"`

	// 站点ID。
	SiteId         *string `json:"site_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowQuotaDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotaDetailsResponse", string(data)}, " ")
}
