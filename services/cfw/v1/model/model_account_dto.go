package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccountDto struct {

	// **参数解释**： 账号列表 **约束限制**： 不涉及 **默认取值**： 不涉及
	Accounts *[]AccountInfo `json:"accounts,omitempty"`
}

func (o AccountDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountDto struct{}"
	}

	return strings.Join([]string{"AccountDto", string(data)}, " ")
}
