package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabaseOmUserInfo 数据库运维账户信息 包含账户状态  账户过期时间（时间戳） **参数解释**： 数据库运维账户信息。包含账户状态  账户过期时间（时间戳）。 **取值范围**： 不涉及。
type DatabaseOmUserInfo struct {

	// **参数解释**： 运维账户状态。 **取值范围**： on、off
	OmUserStatus *string `json:"om_user_status,omitempty"`

	// **参数解释**： 运维账户过期状态。格式是有效的时间戳。 **取值范围**： 不涉及。
	OmUserExpiresTime *int64 `json:"om_user_expires_time,omitempty"`
}

func (o DatabaseOmUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseOmUserInfo struct{}"
	}

	return strings.Join([]string{"DatabaseOmUserInfo", string(data)}, " ")
}
