package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupBatchAddMemberDto struct {

	// **参数解释：** 用户iam_id。 **取值范围：** 字符串长度不少于1，不超过1000。
	IamId *string `json:"iam_id,omitempty"`

	// **参数解释：** 用户userId。 **取值范围：** 字符串长度不少于1，不超过1000。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释：** 用户名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 用户昵称。 **取值范围：** 字符串长度不少于1，不超过1000。
	NickName *string `json:"nick_name,omitempty"`

	// **参数解释：** 租户名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释：** 租户id。 **取值范围：** 字符串长度不少于1，不超过1000。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** 角色id。 **取值范围：** 字符串长度不少于1，不超过1000。
	RepoRoleId *string `json:"repo_role_id,omitempty"`

	// **参数解释：** 项目角色id。 **取值范围：** 字符串长度不少于1，不超过1000。
	ReqRoleId *string `json:"req_role_id,omitempty"`

	// **参数解释：** 角色名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	RepoRoleName *string `json:"repo_role_name,omitempty"`

	// **参数解释：** 项目角色名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	ReqRoleName *string `json:"req_role_name,omitempty"`
}

func (o GroupBatchAddMemberDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupBatchAddMemberDto struct{}"
	}

	return strings.Join([]string{"GroupBatchAddMemberDto", string(data)}, " ")
}
