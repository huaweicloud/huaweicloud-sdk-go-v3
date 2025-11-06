package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApproverBasicDto Get a merge request approver info
type ApproverBasicDto struct {

	// **参数解释：** 用户id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 用户名称。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 用户名。
	Username *string `json:"username,omitempty"`

	// **参数解释：** 中文名。 **取值范围：** 不涉及。
	NameCn *string `json:"name_cn,omitempty"`

	// **参数解释：** 邮件。 **取值范围：** 不涉及。
	Email *string `json:"email,omitempty"`

	// **参数解释：** 状态。 - optional，可选（审核人、检视人）（未审核/检视） - required，必选（审核人、检视人）（未审核/检视） - approve，审核通过 - true，检视通过 - passed，审核/检视 通过，（未审核/检视） - reject，审核拒绝
	State *string `json:"state,omitempty"`

	// **参数解释：** 更新时间 **取值范围：** 不涉及。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 头像链接 **取值范围：** 不涉及。
	AvatarUrl *string `json:"avatar_url,omitempty"`

	// **参数解释：** 昵称 **取值范围：** 不涉及。
	NickName *string `json:"nick_name,omitempty"`

	// **参数解释：** 租户名称。 **取值范围：** 不涉及。
	TenantName *string `json:"tenant_name,omitempty"`

	// **参数解释：** 审核备注。 **取值范围：** 不涉及。
	ApproverComment *string `json:"approver_comment,omitempty"`
}

func (o ApproverBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApproverBasicDto struct{}"
	}

	return strings.Join([]string{"ApproverBasicDto", string(data)}, " ")
}
