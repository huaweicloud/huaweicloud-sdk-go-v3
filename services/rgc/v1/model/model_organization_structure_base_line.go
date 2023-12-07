package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrganizationStructureBaseLine 账号体系基础设置。
type OrganizationStructureBaseLine struct {

	// OU名称。
	OrganizationalUnitName string `json:"organizational_unit_name"`

	OrganizationalUnitType *OrganizationalUnitTypeForSetup `json:"organizational_unit_type"`

	// 账号基本信息。
	Accounts *[]AccountBaseline `json:"accounts,omitempty"`
}

func (o OrganizationStructureBaseLine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationStructureBaseLine struct{}"
	}

	return strings.Join([]string{"OrganizationStructureBaseLine", string(data)}, " ")
}
