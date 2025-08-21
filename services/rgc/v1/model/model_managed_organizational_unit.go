package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ManagedOrganizationalUnit struct {

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

	// 注册OU的更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 错误信息描述。
	Message *string `json:"message,omitempty"`
}

func (o ManagedOrganizationalUnit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ManagedOrganizationalUnit struct{}"
	}

	return strings.Join([]string{"ManagedOrganizationalUnit", string(data)}, " ")
}
