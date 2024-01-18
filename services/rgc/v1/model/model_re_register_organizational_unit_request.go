package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReRegisterOrganizationalUnitRequest Request Object
type ReRegisterOrganizationalUnitRequest struct {

	// 注册OU ID。
	OrganizationUnitId string `json:"organization_unit_id"`
}

func (o ReRegisterOrganizationalUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReRegisterOrganizationalUnitRequest struct{}"
	}

	return strings.Join([]string{"ReRegisterOrganizationalUnitRequest", string(data)}, " ")
}
