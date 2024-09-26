package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpPoliciesResponse Response Object
type ShowHttpPoliciesResponse struct {

	// 防护策略的数量
	Total *int32 `json:"total,omitempty"`

	// 防护策略的具体内容
	Items          *[]ShowHttpPolicyResponseBody `json:"items,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ShowHttpPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpPoliciesResponse", string(data)}, " ")
}
