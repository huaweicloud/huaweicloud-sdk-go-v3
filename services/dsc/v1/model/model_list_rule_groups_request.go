package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRuleGroupsRequest Request Object
type ListRuleGroupsRequest struct {

	// 页码
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRuleGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRuleGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListRuleGroupsRequest", string(data)}, " ")
}
