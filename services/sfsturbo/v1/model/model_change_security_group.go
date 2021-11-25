package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 需要修改的目标安全组ID
type ChangeSecurityGroup struct {
	// 需要修改的目标安全组ID。

	SecurityGroupId string `json:"security_group_id"`
}

func (o ChangeSecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSecurityGroup struct{}"
	}

	return strings.Join([]string{"ChangeSecurityGroup", string(data)}, " ")
}
