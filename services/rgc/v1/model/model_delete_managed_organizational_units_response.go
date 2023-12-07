package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteManagedOrganizationalUnitsResponse Response Object
type DeleteManagedOrganizationalUnitsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteManagedOrganizationalUnitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteManagedOrganizationalUnitsResponse struct{}"
	}

	return strings.Join([]string{"DeleteManagedOrganizationalUnitsResponse", string(data)}, " ")
}
