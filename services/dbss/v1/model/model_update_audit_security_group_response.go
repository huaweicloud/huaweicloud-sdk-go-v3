package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuditSecurityGroupResponse Response Object
type UpdateAuditSecurityGroupResponse struct {

	// 响应状态
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAuditSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuditSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateAuditSecurityGroupResponse", string(data)}, " ")
}
