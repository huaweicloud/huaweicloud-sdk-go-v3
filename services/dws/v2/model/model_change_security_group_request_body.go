package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeSecurityGroupRequestBody 修改集群安全组请求参数
type ChangeSecurityGroupRequestBody struct {

	// 安全组ID数组
	SecurityGroups []string `json:"security_groups"`
}

func (o ChangeSecurityGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeSecurityGroupRequestBody", string(data)}, " ")
}
