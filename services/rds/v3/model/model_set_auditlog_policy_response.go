package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SetAuditlogPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetAuditlogPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAuditlogPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetAuditlogPolicyResponse", string(data)}, " ")
}
