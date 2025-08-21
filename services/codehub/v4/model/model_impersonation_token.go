package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImpersonationToken This feature was introduced in  9.0
type ImpersonationToken struct {

	// **参数解释：** 唯一标示id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 是否撤销。
	Revoked *bool `json:"revoked,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** scope权限。 **取值范围：** 字符串长度不少于0，不超过1000。
	Scopes *[]string `json:"scopes,omitempty"`

	// **参数解释：** 是否可用。
	Active *bool `json:"active,omitempty"`

	// **参数解释：** 过期时间。 **取值范围：** 字符串长度不少于1，不超过1000。
	ExpiresAt *string `json:"expires_at,omitempty"`

	// **参数解释：** 是否为个人token。
	Impersonation *bool `json:"impersonation,omitempty"`

	// **参数解释：** 描述。 **取值范围：** 字符串长度不少于1，不超过1000。
	Description *string `json:"description,omitempty"`
}

func (o ImpersonationToken) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImpersonationToken struct{}"
	}

	return strings.Join([]string{"ImpersonationToken", string(data)}, " ")
}
