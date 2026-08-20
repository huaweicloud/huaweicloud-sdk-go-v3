package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserObject 用户对象, 适用于评审单。
type UserObject struct {

	// 是否关注。
	Watcher *string `json:"watcher,omitempty"`

	// 用户ID。
	UserId *string `json:"user_id,omitempty"`

	// 用户数字ID。
	UserNumId *string `json:"user_num_id,omitempty"`

	// 用户名称。
	UserName *string `json:"user_name,omitempty"`

	// 租户ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 租户名称。
	DomainName *string `json:"domain_name,omitempty"`

	// 昵称。
	NickName *string `json:"nick_name,omitempty"`

	// 角色ID。
	RoleId *string `json:"role_id,omitempty"`

	// 角色名称。
	RoleName *string `json:"role_name,omitempty"`

	// 用户头像。
	ImageId *string `json:"image_id,omitempty"`

	// 区域。
	Region *string `json:"region,omitempty"`

	// 意见。
	Opinion *string `json:"opinion,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 责任人。
	Owner *string `json:"owner,omitempty"`

	// 决策人ID。
	CcbId *string `json:"ccbId,omitempty"`

	// 是否已移出项目。
	HasRemoved *string `json:"has_removed,omitempty"`
}

func (o UserObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserObject struct{}"
	}

	return strings.Join([]string{"UserObject", string(data)}, " ")
}
