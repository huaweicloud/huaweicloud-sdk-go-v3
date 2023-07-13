package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAccessPoliciesRequest Request Object
type BatchDeleteAccessPoliciesRequest struct {
	Body *BatchDeleteAccessPoliciesReq `json:"body,omitempty"`
}

func (o BatchDeleteAccessPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAccessPoliciesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAccessPoliciesRequest", string(data)}, " ")
}
