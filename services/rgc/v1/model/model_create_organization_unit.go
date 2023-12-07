package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrganizationUnit 纳管的OU的基本信息。
type CreateOrganizationUnit struct {

	// 管理账号ID。
	ManageAccountId *string `json:"manage_account_id,omitempty"`

	// OU ID。
	OrganizationUnitId *string `json:"organization_unit_id,omitempty"`

	// OU名称。
	OrganizationUnitName *string `json:"organization_unit_name,omitempty"`

	// OU状态。
	OrganizationUnitStatus *string `json:"organization_unit_status,omitempty"`

	OrganizationUnitType *OrganizationalUnitType `json:"organization_unit_type,omitempty"`

	// 父OU ID。
	ParentOrganizationUnitId *string `json:"parent_organization_unit_id,omitempty"`

	// 父OU名称。
	ParentOrganizationUnitName *string `json:"parent_organization_unit_name,omitempty"`

	// 被创建的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`
}

func (o CreateOrganizationUnit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrganizationUnit struct{}"
	}

	return strings.Join([]string{"CreateOrganizationUnit", string(data)}, " ")
}
