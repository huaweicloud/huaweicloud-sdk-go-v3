package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOrganizationalUnitResponse Response Object
type DeleteOrganizationalUnitResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteOrganizationalUnitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOrganizationalUnitResponse struct{}"
	}

	return strings.Join([]string{"DeleteOrganizationalUnitResponse", string(data)}, " ")
}
