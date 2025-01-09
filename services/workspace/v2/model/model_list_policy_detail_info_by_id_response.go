package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyDetailInfoByIdResponse Response Object
type ListPolicyDetailInfoByIdResponse struct {
	PolicyGroup    *PolicyGroup `json:"policy_group,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListPolicyDetailInfoByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyDetailInfoByIdResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyDetailInfoByIdResponse", string(data)}, " ")
}
