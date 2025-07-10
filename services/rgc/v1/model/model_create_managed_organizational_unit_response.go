package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateManagedOrganizationalUnitResponse Response Object
type CreateManagedOrganizationalUnitResponse struct {

	// 管理纳管账号ID。
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

	// 组织里某个注册OU下的纳管账号被创建的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// Landing Zone版本。
	LandingZoneVersion *string `json:"landing_zone_version,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o CreateManagedOrganizationalUnitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManagedOrganizationalUnitResponse struct{}"
	}

	return strings.Join([]string{"CreateManagedOrganizationalUnitResponse", string(data)}, " ")
}
