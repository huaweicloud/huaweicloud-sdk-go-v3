package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowOrganizationalUnitResponse struct {
	OrganizationalUnit *OrganizationalUnitDto `json:"organizational_unit,omitempty"`
	HttpStatusCode     int                    `json:"-"`
}

func (o ShowOrganizationalUnitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationalUnitResponse struct{}"
	}

	return strings.Join([]string{"ShowOrganizationalUnitResponse", string(data)}, " ")
}
