package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityGroupTagRequest Request Object
type CreateSecurityGroupTagRequest struct {

	// 安全组资源ID
	SecurityGroupId string `json:"security_group_id"`

	Body *CreateSecurityGroupTagRequestBody `json:"body,omitempty"`
}

func (o CreateSecurityGroupTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupTagRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupTagRequest", string(data)}, " ")
}
