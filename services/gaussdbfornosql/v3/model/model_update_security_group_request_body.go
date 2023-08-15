package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSecurityGroupRequestBody struct {

	// 新的安全组ID。
	SecurityGroupId string `json:"security_group_id"`
}

func (o UpdateSecurityGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSecurityGroupRequestBody", string(data)}, " ")
}
