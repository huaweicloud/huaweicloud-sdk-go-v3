package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRiskOperationPolicyResponse Response Object
type SetRiskOperationPolicyResponse struct {

	// 状态  - SUCCESS:成功  - FAILED:失败
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetRiskOperationPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRiskOperationPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetRiskOperationPolicyResponse", string(data)}, " ")
}
