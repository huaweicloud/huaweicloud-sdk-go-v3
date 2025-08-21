package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrganizationalUnit 注册OU的基本信息。
type CreateOrganizationalUnit struct {

	// 管理员账号ID。
	ManageAccountId *string `json:"manage_account_id,omitempty"`

	// 注册OU ID。
	OrganizationalUnitId *string `json:"organizational_unit_id,omitempty"`

	// 注册OU名称。
	OrganizationalUnitName *string `json:"organizational_unit_name,omitempty"`

	// 注册OU状态。
	OrganizationalUnitStatus *string `json:"organizational_unit_status,omitempty"`

	OrganizationalUnitType *OrganizationalUnitType `json:"organizational_unit_type,omitempty"`

	// 父注册OU ID。
	ParentOrganizationalUnitId *string `json:"parent_organizational_unit_id,omitempty"`

	// 父注册OU名称。
	ParentOrganizationalUnitName *string `json:"parent_organizational_unit_name,omitempty"`

	// 注册OU的创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 注册OU的Landing Zone版本。
	LandingZoneVersion *string `json:"landing_zone_version,omitempty"`
}

func (o CreateOrganizationalUnit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrganizationalUnit struct{}"
	}

	return strings.Join([]string{"CreateOrganizationalUnit", string(data)}, " ")
}
