package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAuditlogPolicyResponse Response Object
type SetAuditlogPolicyResponse struct {

	// 设置审计日志策略的操作结果。 - COMPLETED：已完成。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetAuditlogPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAuditlogPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetAuditlogPolicyResponse", string(data)}, " ")
}
