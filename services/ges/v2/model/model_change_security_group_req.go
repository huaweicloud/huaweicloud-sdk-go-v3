package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeSecurityGroupReq struct {

	// 目标安全组的ID。
	SecurityGroupId string `json:"security_group_id"`
}

func (o ChangeSecurityGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSecurityGroupReq struct{}"
	}

	return strings.Join([]string{"ChangeSecurityGroupReq", string(data)}, " ")
}
