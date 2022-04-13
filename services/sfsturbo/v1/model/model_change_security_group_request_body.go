package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// change_security_group对象
type ChangeSecurityGroupRequestBody struct {
	ChangeSecurityGroup *ChangeSecurityGroup `json:"change_security_group"`
}

func (o ChangeSecurityGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeSecurityGroupRequestBody", string(data)}, " ")
}
