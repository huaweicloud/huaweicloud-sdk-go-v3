package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto update body Object
type ModifyInstanceSecurityGroupReq struct {
	// 安全组ID，默认值为原安全组ID，可根据需要判断是否修改安全组ID

	SecurityGroupId string `json:"security_group_id"`
}

func (o ModifyInstanceSecurityGroupReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInstanceSecurityGroupReq struct{}"
	}

	return strings.Join([]string{"ModifyInstanceSecurityGroupReq", string(data)}, " ")
}
