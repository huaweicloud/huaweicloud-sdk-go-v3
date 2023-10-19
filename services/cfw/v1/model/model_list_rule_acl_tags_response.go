package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRuleAclTagsResponse Response Object
type ListRuleAclTagsResponse struct {
	Data           *HttpGetAclTagResponseData `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListRuleAclTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleAclTagsResponse struct{}"
	}

	return strings.Join([]string{"ListRuleAclTagsResponse", string(data)}, " ")
}
