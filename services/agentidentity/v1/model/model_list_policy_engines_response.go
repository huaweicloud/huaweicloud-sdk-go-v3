package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyEnginesResponse Response Object
type ListPolicyEnginesResponse struct {
	PolicyEngines *[]PolicyEngineSummary `json:"policy_engines,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPolicyEnginesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyEnginesResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyEnginesResponse", string(data)}, " ")
}
