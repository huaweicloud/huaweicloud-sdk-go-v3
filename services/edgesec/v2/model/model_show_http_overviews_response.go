package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpOverviewsResponse Response Object
type ShowHttpOverviewsResponse struct {
	Domain *DomainClassificationItem `json:"domain,omitempty"`

	AttackType *AttackTypeClassificationItem `json:"attack_type,omitempty"`

	Ip *IpClassificationItem `json:"ip,omitempty"`

	Url *UrlClassificationItem `json:"url,omitempty"`

	Geo            *GeoClassificationItem `json:"geo,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowHttpOverviewsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpOverviewsResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpOverviewsResponse", string(data)}, " ")
}
