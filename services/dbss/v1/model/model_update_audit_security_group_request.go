package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuditSecurityGroupRequest Request Object
type UpdateAuditSecurityGroupRequest struct {
	Body *SecurityGroupRequest `json:"body,omitempty"`
}

func (o UpdateAuditSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuditSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateAuditSecurityGroupRequest", string(data)}, " ")
}
