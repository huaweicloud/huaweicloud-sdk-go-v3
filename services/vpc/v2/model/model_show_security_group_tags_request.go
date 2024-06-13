package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityGroupTagsRequest Request Object
type ShowSecurityGroupTagsRequest struct {

	// 安全组资源ID
	SecurityGroupId string `json:"security_group_id"`
}

func (o ShowSecurityGroupTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityGroupTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityGroupTagsRequest", string(data)}, " ")
}
