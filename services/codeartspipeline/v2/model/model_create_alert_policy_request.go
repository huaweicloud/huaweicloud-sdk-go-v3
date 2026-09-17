package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlertPolicyRequest Request Object
type CreateAlertPolicyRequest struct {

	// 租户ID
	TenantId string `json:"tenant_id"`

	Body *AlertPolicyDto `json:"body,omitempty"`
}

func (o CreateAlertPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlertPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateAlertPolicyRequest", string(data)}, " ")
}
