package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuditSecurityGroupNewRequest Request Object
type UpdateAuditSecurityGroupNewRequest struct {
	Body *SecurityGroupRequest `json:"body,omitempty"`
}

func (o UpdateAuditSecurityGroupNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuditSecurityGroupNewRequest struct{}"
	}

	return strings.Join([]string{"UpdateAuditSecurityGroupNewRequest", string(data)}, " ")
}
