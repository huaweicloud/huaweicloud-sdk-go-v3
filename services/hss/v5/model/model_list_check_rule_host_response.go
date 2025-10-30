package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCheckRuleHostResponse Response Object
type ListCheckRuleHostResponse struct {

	// **参数解释** 单个基线检查项影响的主机总量 **取值范围** 取值0-2147483647
	TotalNum *int64 `json:"total_num,omitempty"`

	// **参数解释** 单个基线检查项受影响的主机列表 **取值范围** 最少0条，最多2147483647条
	DataList       *[]SecurityCheckRuleHostResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ListCheckRuleHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCheckRuleHostResponse struct{}"
	}

	return strings.Join([]string{"ListCheckRuleHostResponse", string(data)}, " ")
}
