package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryUserDo struct {

	// **参数解释**: 账号。 **取值范围**: 不涉及。
	Username *string `json:"username,omitempty"`

	// **参数解释**: 密码。 **取值范围**: 不涉及。
	Password *string `json:"password,omitempty"`
}

func (o RepositoryUserDo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryUserDo struct{}"
	}

	return strings.Join([]string{"RepositoryUserDo", string(data)}, " ")
}
