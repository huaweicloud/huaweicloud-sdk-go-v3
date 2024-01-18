package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ManagedOrganizationUnit struct {

	// 管理纳管账号ID。
	ManageAccountId *string `json:"manage_account_id,omitempty"`

	// 注册OU ID。
	OrganizationUnitId *string `json:"organization_unit_id,omitempty"`

	// 注册OU名称。
	OrganizationUnitName *string `json:"organization_unit_name,omitempty"`

	// 注册OU状态。
	OrganizationUnitStatus *string `json:"organization_unit_status,omitempty"`

	OrganizationUnitType *OrganizationalUnitType `json:"organization_unit_type,omitempty"`

	// 父注册OU ID。
	ParentOrganizationUnitId *string `json:"parent_organization_unit_id,omitempty"`

	// 父注册OU名称。
	ParentOrganizationUnitName *string `json:"parent_organization_unit_name,omitempty"`

	// 组织里某个注册OU下的纳管账号被创建的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// Landing Zone版本。
	LandingZoneVersion *string `json:"landing_zone_version,omitempty"`

	// 组织里某个注册OU下的纳管账号最后一次更新的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 错误信息描述。
	Message *string `json:"message,omitempty"`
}

func (o ManagedOrganizationUnit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ManagedOrganizationUnit struct{}"
	}

	return strings.Join([]string{"ManagedOrganizationUnit", string(data)}, " ")
}
