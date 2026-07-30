package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiPoliciesResponse Response Object
type ListAiPoliciesResponse struct {

	// **参数解释**: AI策略列表 **取值范围**: 取值0-200个对象
	DataList       *[]AiPolicyInfo `json:"data_list,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListAiPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListAiPoliciesResponse", string(data)}, " ")
}
