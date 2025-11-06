package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BlameDto struct {
	Commit *RepositoryCommitDto `json:"commit,omitempty"`

	// **参数解释：** 头像链接 **取值范围：** 不涉及。
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// 行信息
	Lines *[]LineContentDto `json:"lines,omitempty"`

	// **参数解释：** 昵称 **取值范围：** 不涉及。
	NickName *string `json:"nick_name,omitempty"`

	// **参数解释：** 租户名称。 **取值范围：** 不涉及。
	TenantName *string `json:"tenant_name,omitempty"`

	// **参数解释：** 用户名。 **取值范围：** 不涉及。
	UserName *string `json:"user_name,omitempty"`
}

func (o BlameDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlameDto struct{}"
	}

	return strings.Join([]string{"BlameDto", string(data)}, " ")
}
