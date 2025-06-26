package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabaseUser **参数解释**： 用户基本信息。 **取值范围**： 不涉及。
type DatabaseUser struct {

	// **参数解释**： 用户名。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 是否可以登录。 **取值范围**： 不涉及。
	Login *bool `json:"login,omitempty"`

	// **参数解释**： 用户类型。 **取值范围**： COMMON：表示普通数据库用户。 IAM：表示IAM同步的数据库用户。 OneAccess: 表示OneAccess用户。
	UserType *string `json:"user_type,omitempty"`
}

func (o DatabaseUser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseUser struct{}"
	}

	return strings.Join([]string{"DatabaseUser", string(data)}, " ")
}
