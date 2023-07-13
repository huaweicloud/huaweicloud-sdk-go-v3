package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrganizationalUnitResponse Response Object
type CreateOrganizationalUnitResponse struct {
	OrganizationalUnit *OrganizationalUnitDto `json:"organizational_unit,omitempty"`
	HttpStatusCode     int                    `json:"-"`
}

func (o CreateOrganizationalUnitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrganizationalUnitResponse struct{}"
	}

	return strings.Join([]string{"CreateOrganizationalUnitResponse", string(data)}, " ")
}
