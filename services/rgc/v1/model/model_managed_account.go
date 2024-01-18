package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ManagedAccount 注册OU信息。
type ManagedAccount struct {

	// Landing Zone版本。
	LandingZoneVersion *string `json:"landing_zone_version,omitempty"`

	// 管理纳管账号ID。
	ManageAccountId *string `json:"manage_account_id,omitempty"`

	// 纳管账号ID。
	AccountId *string `json:"account_id,omitempty"`

	// 纳管账号名称。
	AccountName *string `json:"account_name,omitempty"`

	// 纳管账号类型。
	AccountType *string `json:"account_type,omitempty"`

	// 纳管账号的创建来源，包括CUSTOM和RGC。
	Owner *string `json:"owner,omitempty"`

	// 纳管账号状态。
	State *string `json:"state,omitempty"`

	// 错误状态描述信息。
	Message *string `json:"message,omitempty"`

	// 父注册OU ID。
	ParentOrganizationUnitId *string `json:"parent_organization_unit_id,omitempty"`

	// 父注册OU名称。
	ParentOrganizationUnitName *string `json:"parent_organization_unit_name,omitempty"`

	// Identity Center用户名。
	IdentityStoreUserName *string `json:"identity_store_user_name,omitempty"`

	// 模板ID。
	BlueprintProductId *string `json:"blueprint_product_id,omitempty"`

	// 模板版本。
	BlueprintProductVersion *string `json:"blueprint_product_version,omitempty"`

	// 模板部署状态，包括失败, 完成, 进行中。
	BlueprintStatus *string `json:"blueprint_status,omitempty"`

	// 区域信息。
	Regions *[]RegionManagedList `json:"regions,omitempty"`

	// 组织里某个注册OU下的纳管账号被创建的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 组织里某个注册OU下的纳管账号最后一次更新的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o ManagedAccount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ManagedAccount struct{}"
	}

	return strings.Join([]string{"ManagedAccount", string(data)}, " ")
}
