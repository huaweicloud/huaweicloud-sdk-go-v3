package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRestrictionOfInstanceV2Response Response Object
type ShowRestrictionOfInstanceV2Response struct {

	// 受限的IP网段的CIDR列表。
	RestrictCidrs *[]string `json:"restrict_cidrs,omitempty"`

	// 资源租户的IP网段的CIDR。
	ResourceSubnetCidr *string `json:"resource_subnet_cidr,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o ShowRestrictionOfInstanceV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRestrictionOfInstanceV2Response struct{}"
	}

	return strings.Join([]string{"ShowRestrictionOfInstanceV2Response", string(data)}, " ")
}
