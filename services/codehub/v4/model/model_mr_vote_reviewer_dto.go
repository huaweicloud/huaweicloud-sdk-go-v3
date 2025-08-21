package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MrVoteReviewerDto **参数解释：** 合并请求检视用户详细信息
type MrVoteReviewerDto struct {

	// **参数解释：** 用户id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 用户名称。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 用户名。
	Username *string `json:"username,omitempty"`

	// 用户状态
	State *string `json:"state,omitempty"`

	// 服务级权限状态 0：停用 1：启用
	ServiceLicenseStatus *int32 `json:"service_license_status,omitempty"`

	// 用户头像url
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// 用户头像路径
	AvatarPath *string `json:"avatar_path,omitempty"`

	// 用户邮箱
	Email *string `json:"email,omitempty"`

	// 用户中文名称
	NameCn *string `json:"name_cn,omitempty"`

	// 用户个人首页
	WebUrl *string `json:"web_url,omitempty"`

	// 用户昵称
	NickName *string `json:"nick_name,omitempty"`

	// 租户名称
	TenantName *string `json:"tenant_name,omitempty"`

	// **参数解释：** 部分查询接口校验到传参里的用户权限不足或不存在时，返回该用户但该字段不为空用于提示。
	ErrorMessage *string `json:"error_message,omitempty"`

	// **参数解释：** 是否为committer。
	IsCommitter *bool `json:"is_committer,omitempty"`

	// **参数解释：** 是否为认证committer。
	IsVerified *bool `json:"is_verified,omitempty"`

	// **参数解释：** 是否有相关权限。
	HasPermission *bool `json:"has_permission,omitempty"`
}

func (o MrVoteReviewerDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MrVoteReviewerDto struct{}"
	}

	return strings.Join([]string{"MrVoteReviewerDto", string(data)}, " ")
}
