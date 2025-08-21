package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRepositoryMemberResponseDto struct {

	// **参数解释：** 用户iamId **约束限制：** 不涉及。
	UserIamId *string `json:"user_iam_id,omitempty"`

	// **参数解释：** 用户名称。 **约束限制：** 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释：** 租户名称。 **约束限制：** - 不涉及。
	TenantName *string `json:"tenant_name,omitempty"`

	// **参数解释：** 用户昵称。 **约束限制：** - 不涉及。
	UserNickName *string `json:"user_nick_name,omitempty"`

	// **参数解释：** 角色id。 **约束限制：** **约束限制：** - success，添加成功。 - error，添加失败。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 成员添加结果信息。 **约束限制：** - 不涉及。
	Message *string `json:"message,omitempty"`
}

func (o CreateRepositoryMemberResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepositoryMemberResponseDto struct{}"
	}

	return strings.Join([]string{"CreateRepositoryMemberResponseDto", string(data)}, " ")
}
