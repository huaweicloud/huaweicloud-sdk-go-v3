package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRestrictedRuleRequest Request Object
type ListRestrictedRuleRequest struct {

	// 查询的偏移量，默认值0。
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询的大小[1-100]，默认值10。
	Limit *int32 `json:"limit,omitempty"`

	// 应用规则名称。
	Name *string `json:"name,omitempty"`
}

func (o ListRestrictedRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestrictedRuleRequest struct{}"
	}

	return strings.Join([]string{"ListRestrictedRuleRequest", string(data)}, " ")
}
