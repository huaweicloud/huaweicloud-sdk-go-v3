package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomRuleConfigsResponse Response Object
type ListCustomRuleConfigsResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 自定义规则列表 **取值范围**: 取值0-200个ListCustomRuleConfigResponse对象
	DataList       *[]ListCustomRuleConfigResponse `json:"data_list,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListCustomRuleConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomRuleConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListCustomRuleConfigsResponse", string(data)}, " ")
}
