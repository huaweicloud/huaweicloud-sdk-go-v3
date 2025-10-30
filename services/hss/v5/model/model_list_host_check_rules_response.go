package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostCheckRulesResponse Response Object
type ListHostCheckRulesResponse struct {

	// **参数解释** 单个基线检查项影响的主机总量 **约束限制** 不涉及 **取值范围** 不涉及 **默认取值** 不涉及
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 单个主机涉及的检查项列表 **约束限制** 不涉及 **取值范围** 不涉及 **默认取值** 不涉及
	DataList       *[]HostCheckRulesResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListHostCheckRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostCheckRulesResponse struct{}"
	}

	return strings.Join([]string{"ListHostCheckRulesResponse", string(data)}, " ")
}
