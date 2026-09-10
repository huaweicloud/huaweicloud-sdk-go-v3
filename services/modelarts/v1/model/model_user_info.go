package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserInfo struct {

	// **参数解释**：授权对象名称。 **取值范围**：不涉及。
	UserName *string `json:"userName,omitempty"`
}

func (o UserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserInfo struct{}"
	}

	return strings.Join([]string{"UserInfo", string(data)}, " ")
}
