package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgencyResponse Response Object
type ShowAgencyResponse struct {

	// 当前存在的委托权限列表。
	ExistingAgencies *[]AgencyDto `json:"existing_agencies,omitempty"`

	// 仍需要的委托权限列表。
	WantedAgencies *[]AgencyDto `json:"wanted_agencies,omitempty"`

	// 冗余的委托权限列表。
	RedundantAgencies *[]AgencyDto `json:"redundant_agencies,omitempty"`
	HttpStatusCode    int          `json:"-"`
}

func (o ShowAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyResponse struct{}"
	}

	return strings.Join([]string{"ShowAgencyResponse", string(data)}, " ")
}
