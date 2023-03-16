package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteOrganizationalUnitRequest struct {

	// 与组织单元关联的唯一标识符（ID）。
	OrganizationalUnitId string `json:"organizational_unit_id"`
}

func (o DeleteOrganizationalUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOrganizationalUnitRequest struct{}"
	}

	return strings.Join([]string{"DeleteOrganizationalUnitRequest", string(data)}, " ")
}
