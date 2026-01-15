package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuditSecurityGroupNewResponse Response Object
type UpdateAuditSecurityGroupNewResponse struct {

	// 操作结果  - success：成功  - failed：失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAuditSecurityGroupNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuditSecurityGroupNewResponse struct{}"
	}

	return strings.Join([]string{"UpdateAuditSecurityGroupNewResponse", string(data)}, " ")
}
