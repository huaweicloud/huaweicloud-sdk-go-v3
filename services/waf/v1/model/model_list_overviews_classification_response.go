package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOverviewsClassificationResponse Response Object
type ListOverviewsClassificationResponse struct {
	Domain *DomainClassificationItem `json:"domain,omitempty"`

	AttackType *AttackTypeClassificationItem `json:"attack_type,omitempty"`

	Ip *IpClassificationItem `json:"ip,omitempty"`

	Url *UrlClassificationItem `json:"url,omitempty"`

	Geo            *GeoClassificationItem `json:"geo,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListOverviewsClassificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOverviewsClassificationResponse struct{}"
	}

	return strings.Join([]string{"ListOverviewsClassificationResponse", string(data)}, " ")
}
