package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrganizationStructureBaseLineRsp 纳管账号体系基础设置。
type OrganizationStructureBaseLineRsp struct {

	// 注册OU名称。
	OrganizationalUnitName *string `json:"organizational_unit_name,omitempty"`

	OrganizationalUnitType *OrganizationalUnitTypeForSetup `json:"organizational_unit_type"`

	// 纳管账号基本信息。
	Accounts *[]AccountBaselineRsp `json:"accounts,omitempty"`
}

func (o OrganizationStructureBaseLineRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrganizationStructureBaseLineRsp struct{}"
	}

	return strings.Join([]string{"OrganizationStructureBaseLineRsp", string(data)}, " ")
}
