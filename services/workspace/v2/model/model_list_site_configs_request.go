package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSiteConfigsRequest Request Object
type ListSiteConfigsRequest struct {

	// 可用区
	AvailabilityZoneId *string `json:"availability_zone_id,omitempty"`

	// 站点类型，支持CENTER、IES。
	SiteType *string `json:"site_type,omitempty"`
}

func (o ListSiteConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSiteConfigsRequest struct{}"
	}

	return strings.Join([]string{"ListSiteConfigsRequest", string(data)}, " ")
}
