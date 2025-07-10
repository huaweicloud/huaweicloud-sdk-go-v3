package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotaDetailsRequest Request Object
type ShowQuotaDetailsRequest struct {

	// 站点ID。
	SiteId *string `json:"site_id,omitempty"`

	// 可用分区code。
	AzCode *string `json:"az_code,omitempty"`
}

func (o ShowQuotaDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotaDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowQuotaDetailsRequest", string(data)}, " ")
}
