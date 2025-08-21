package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContributorDto struct {

	// **参数解释：** 用户名。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 邮箱。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Email *string `json:"email,omitempty"`

	// **参数解释：** 提交数量。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Commits *int32 `json:"commits,omitempty"`

	// **参数解释：** 昵称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	NickName *string `json:"nick_name,omitempty"`

	// **参数解释：** 租户名。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	TenantName *string `json:"tenant_name,omitempty"`

	// **参数解释：** 用户名。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	UserName *string `json:"user_name,omitempty"`
}

func (o ContributorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContributorDto struct{}"
	}

	return strings.Join([]string{"ContributorDto", string(data)}, " ")
}
