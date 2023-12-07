package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateManagedAccountRequest 创建账号的基本信息。
type CreateManagedAccountRequest struct {

	// 账号名。
	AccountName string `json:"account_name"`

	// 账号邮箱。
	AccountEmail string `json:"account_email"`

	// 手机号码。
	Phone *string `json:"phone,omitempty"`

	// Identity Center用户名。
	IdentityStoreUserName string `json:"identity_store_user_name"`

	// Identity Center邮箱。
	IdentityStoreEmail string `json:"identity_store_email"`

	// 父OU ID。
	ParentOrganizationUnitId string `json:"parent_organization_unit_id"`

	// 父OU名称。
	ParentOrganizationUnitName string `json:"parent_organization_unit_name"`

	Blueprint *Blueprint `json:"blueprint,omitempty"`
}

func (o CreateManagedAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManagedAccountRequest struct{}"
	}

	return strings.Join([]string{"CreateManagedAccountRequest", string(data)}, " ")
}
