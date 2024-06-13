package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityGroupTagRequest Request Object
type DeleteSecurityGroupTagRequest struct {

	// 功能说明：键值
	Key string `json:"key"`

	// 安全组资源ID
	SecurityGroupId string `json:"security_group_id"`
}

func (o DeleteSecurityGroupTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityGroupTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityGroupTagRequest", string(data)}, " ")
}
