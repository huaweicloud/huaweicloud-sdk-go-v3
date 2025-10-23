package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyResponse Response Object
type ListPolicyResponse struct {

	// count
	Count *int32 `json:"count,omitempty"`

	// bccPolicies
	BccPolicies    *[]BccPolicy `json:"bcc_policies,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyResponse", string(data)}, " ")
}
