package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterOrganizationalUnitRequest Request Object
type RegisterOrganizationalUnitRequest struct {

	// 注册OU ID。
	OrganizationalUnitId string `json:"organizational_unit_id"`
}

func (o RegisterOrganizationalUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterOrganizationalUnitRequest struct{}"
	}

	return strings.Join([]string{"RegisterOrganizationalUnitRequest", string(data)}, " ")
}
