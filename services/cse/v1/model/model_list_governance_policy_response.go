package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGovernancePolicyResponse Response Object
type ListGovernancePolicyResponse struct {

	// 查询治理策略列表响应结构体。
	Body           *[]GovPolicyDetail `json:"body,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListGovernancePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGovernancePolicyResponse struct{}"
	}

	return strings.Join([]string{"ListGovernancePolicyResponse", string(data)}, " ")
}
