package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ManagedOrganizationUnit struct {

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

	// 最后一次更新的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o ManagedOrganizationUnit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ManagedOrganizationUnit struct{}"
	}

	return strings.Join([]string{"ManagedOrganizationUnit", string(data)}, " ")
}
