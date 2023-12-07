package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ManagedAccount 纳管的OU信息。
type ManagedAccount struct {

	// 手机号码。
	Phone *string `json:"phone,omitempty"`

	// 管理账号ID。
	ManageAccountId *string `json:"manage_account_id,omitempty"`

	// 账号ID。
	AccountId *string `json:"account_id,omitempty"`

	// 账号名称。
	AccountName *string `json:"account_name,omitempty"`

	// 账号email。
	AccountEmail *string `json:"account_email,omitempty"`

	// 账号类型。
	AccountType *string `json:"account_type,omitempty"`

	// 账号的创建来源，包括CUSTOM和RGC。
	Owner *string `json:"owner,omitempty"`

	// 账号状态。
	State *string `json:"state,omitempty"`

	// 错误状态描述信息。
	Message *string `json:"message,omitempty"`

	// 父OU ID。
	ParentOrganizationUnitId *string `json:"parent_organization_unit_id,omitempty"`

	// 父OU名称。
	ParentOrganizationUnitName *string `json:"parent_organization_unit_name,omitempty"`

	// Identity Center用户名。
	IdentityStoreUserName *string `json:"identity_store_user_name,omitempty"`

	// Identity Center邮箱。
	IdentityStoreEmailName *string `json:"identity_store_email_name,omitempty"`

	// 蓝图ID。
	BlueprintProductId *string `json:"blueprint_product_id,omitempty"`

	// 蓝图版本。
	BlueprintProductVersion *string `json:"blueprint_product_version,omitempty"`

	// 蓝图部署状态，包括error, active, in_progress。
	BlueprintStatus *string `json:"blueprint_status,omitempty"`

	// region信息。
	Regions *[]RegionManagedList `json:"regions,omitempty"`

	// 被创建的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 最后一次更新的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o ManagedAccount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ManagedAccount struct{}"
	}

	return strings.Join([]string{"ManagedAccount", string(data)}, " ")
}
