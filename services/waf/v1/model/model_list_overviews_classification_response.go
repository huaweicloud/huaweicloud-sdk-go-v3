package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListOverviewsClassificationResponse struct {
	Domain *DomainClassificationItem `json:"domain,omitempty" xml:"domain"`

	AttackType *AttackTypeClassificationItem `json:"attack_type,omitempty" xml:"attack_type"`

	Ip *IpClassificationItem `json:"ip,omitempty" xml:"ip"`

	Url *UrlClassificationItem `json:"url,omitempty" xml:"url"`

	Geo            *GeoClassificationItem `json:"geo,omitempty" xml:"geo"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListOverviewsClassificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOverviewsClassificationResponse struct{}"
	}

	return strings.Join([]string{"ListOverviewsClassificationResponse", string(data)}, " ")
}
