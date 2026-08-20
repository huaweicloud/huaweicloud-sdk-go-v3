package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessInstanceResponseResultCurrOwner 当前责任人
type ProcessInstanceResponseResultCurrOwner struct {

	// 观察者
	Watcher *string `json:"watcher,omitempty"`

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 用户数字id
	UserNumId *string `json:"user_num_id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 租户id
	DomainId *string `json:"domain_id,omitempty"`

	// 租户名
	DomainName *string `json:"domain_name,omitempty"`

	// 昵称
	NickName *string `json:"nick_name,omitempty"`

	// 角色id
	RoleId *string `json:"role_id,omitempty"`

	// 角色名
	RoleName *string `json:"role_name,omitempty"`

	// 用户头像
	ImageId *string `json:"image_id,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 意见
	Opinion *string `json:"opinion,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 责任人
	Owner *string `json:"owner,omitempty"`

	// 评审id
	CcbId *string `json:"ccbId,omitempty"`

	// 是否已移出项目
	HasRemoved *string `json:"has_removed,omitempty"`
}

func (o ProcessInstanceResponseResultCurrOwner) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessInstanceResponseResultCurrOwner struct{}"
	}

	return strings.Join([]string{"ProcessInstanceResponseResultCurrOwner", string(data)}, " ")
}
