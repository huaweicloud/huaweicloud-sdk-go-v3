package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserDetailResponse Response Object
type ShowUserDetailResponse struct {

	// 用户id。
	UserId *string `json:"user_id,omitempty"`

	// 用户所属组织。
	OrgId *string `json:"org_id,omitempty"`

	// 用户名。
	UserName *string `json:"user_name,omitempty"`

	// 姓名。
	Name *string `json:"name,omitempty"`

	// 手机号。
	Mobile *string `json:"mobile,omitempty"`

	// 邮箱。
	Email *string `json:"email,omitempty"`

	// 首次登录是否强制修改密码。
	PwdMustModify *bool `json:"pwd_must_modify,omitempty"`

	// 密码修改时间
	PwdChangeAt *string `json:"pwd_change_at,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 最后一次修改时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 是否禁用。
	Disabled *bool `json:"disabled,omitempty"`

	// 可信等级。
	Grade *int32 `json:"grade,omitempty"`

	// 是否锁定。
	Locked *string `json:"locked,omitempty"`

	// 自定义扩展属性。
	Extension *interface{} `json:"extension,omitempty"`

	// 用户组织关系集合。
	UserOrgRelationList *[]UserOrgRelationListResult `json:"user_org_relation_list,omitempty"`

	// 所属租户ID
	DomainId       *string `json:"domain_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowUserDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowUserDetailResponse", string(data)}, " ")
}
