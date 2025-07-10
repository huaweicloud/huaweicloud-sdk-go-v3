package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteManagedOrganizationalUnitsRequest Request Object
type DeleteManagedOrganizationalUnitsRequest struct {

	// 注册OU ID。
	ManagedOrganizationalUnitId string `json:"managed_organizational_unit_id"`
}

func (o DeleteManagedOrganizationalUnitsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteManagedOrganizationalUnitsRequest struct{}"
	}

	return strings.Join([]string{"DeleteManagedOrganizationalUnitsRequest", string(data)}, " ")
}
