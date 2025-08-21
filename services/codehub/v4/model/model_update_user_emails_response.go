package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserEmailsResponse Response Object
type UpdateUserEmailsResponse struct {

	// **参数解释：** 用户id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 用户名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 用户长id。 **取值范围：** 字符串长度不少于1，不超过1000。
	Username *string `json:"username,omitempty"`

	// **参数解释：** 用户状态。 **取值范围：** 字符串长度不少于1，不超过1000。
	State *string `json:"state,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 最后活跃时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	LastActivityOn *string `json:"last_activity_on,omitempty"`

	// **参数解释：** 提交邮箱。 **取值范围：** 字符串长度不少于1，不超过1000。
	CommitEmail    *string `json:"commit_email,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateUserEmailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserEmailsResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserEmailsResponse", string(data)}, " ")
}
