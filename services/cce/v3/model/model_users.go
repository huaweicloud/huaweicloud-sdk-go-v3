package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Users struct {

	// **参数解释**： 名称 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： user
	Name *string `json:"name,omitempty"`

	User *User `json:"user,omitempty"`
}

func (o Users) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Users struct{}"
	}

	return strings.Join([]string{"Users", string(data)}, " ")
}
