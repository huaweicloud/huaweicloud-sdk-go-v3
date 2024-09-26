package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateManagedAccountRequest 创建账号的基本信息。
type CreateManagedAccountRequest struct {

	// 纳管账号名。
	AccountName string `json:"account_name"`

	// 纳管账号邮箱。
	AccountEmail *string `json:"account_email,omitempty"`

	// 手机号码。
	Phone *string `json:"phone,omitempty"`

	// Identity Center用户名。
	IdentityStoreUserName *string `json:"identity_store_user_name,omitempty"`

	// Identity Center邮箱。
	IdentityStoreEmail *string `json:"identity_store_email,omitempty"`

	// 父注册OU ID。
	ParentOrganizationalUnitId string `json:"parent_organizational_unit_id"`

	// 父注册OU名称。
	ParentOrganizationalUnitName string `json:"parent_organizational_unit_name"`

	Blueprint *Blueprint `json:"blueprint,omitempty"`
}

func (o CreateManagedAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManagedAccountRequest struct{}"
	}

	return strings.Join([]string{"CreateManagedAccountRequest", string(data)}, " ")
}
