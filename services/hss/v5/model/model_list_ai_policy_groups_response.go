package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiPolicyGroupsResponse Response Object
type ListAiPolicyGroupsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: AI策略组列表 **取值范围**: 不涉及
	DataList       *[]AiPolicyGroupInfo `json:"data_list,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListAiPolicyGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiPolicyGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListAiPolicyGroupsResponse", string(data)}, " ")
}
