package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPoliciesResponse Response Object
type ListPoliciesResponse struct {
	Policies *[]PolicySummary `json:"policies,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListPoliciesResponse", string(data)}, " ")
}
