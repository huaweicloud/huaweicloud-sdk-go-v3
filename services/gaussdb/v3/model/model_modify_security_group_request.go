package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifySecurityGroupRequest struct {

	// 安全组ID。
	SecurityGroupId string `json:"security_group_id"`
}

func (o ModifySecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"ModifySecurityGroupRequest", string(data)}, " ")
}
