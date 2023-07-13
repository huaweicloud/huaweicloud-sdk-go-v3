package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOrganizationalUnitResponse Response Object
type UpdateOrganizationalUnitResponse struct {
	OrganizationalUnit *OrganizationalUnitDto `json:"organizational_unit,omitempty"`
	HttpStatusCode     int                    `json:"-"`
}

func (o UpdateOrganizationalUnitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOrganizationalUnitResponse struct{}"
	}

	return strings.Join([]string{"UpdateOrganizationalUnitResponse", string(data)}, " ")
}
