package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteSecurityDynamicMaskingPoliciesResponse Response Object
type BatchDeleteSecurityDynamicMaskingPoliciesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteSecurityDynamicMaskingPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSecurityDynamicMaskingPoliciesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteSecurityDynamicMaskingPoliciesResponse", string(data)}, " ")
}
