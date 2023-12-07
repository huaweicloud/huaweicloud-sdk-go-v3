package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityDataClassificationRulesResponse Response Object
type ListSecurityDataClassificationRulesResponse struct {

	// 查询到的所有数据识别规则
	Content *[]DataClassificationRuleQueryDto `json:"content,omitempty"`

	// 数据识别规则总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSecurityDataClassificationRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityDataClassificationRulesResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityDataClassificationRulesResponse", string(data)}, " ")
}
