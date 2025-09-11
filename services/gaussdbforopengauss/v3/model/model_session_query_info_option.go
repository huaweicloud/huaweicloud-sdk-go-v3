package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SessionQueryInfoOption **参数解释**： 获取实时会话筛选条件。 **约束限制**： 不涉及。
type SessionQueryInfoOption struct {

	// **参数解释**： 数据库名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DatabaseName *string `json:"database_name,omitempty"`

	// **参数解释**： 客户端IP。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClientIp *string `json:"client_ip,omitempty"`

	// **参数解释**： 用户名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	UserName *string `json:"user_name,omitempty"`
}

func (o SessionQueryInfoOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionQueryInfoOption struct{}"
	}

	return strings.Join([]string{"SessionQueryInfoOption", string(data)}, " ")
}
