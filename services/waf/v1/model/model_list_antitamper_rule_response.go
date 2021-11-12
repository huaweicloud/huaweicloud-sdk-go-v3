package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAntitamperRuleResponse struct {
	// 总条数

	Total *int32 `json:"total,omitempty"`
	// 规则

	Items          *[]AntiTamperRuleResponseBody `json:"items,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListAntitamperRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAntitamperRuleResponse struct{}"
	}

	return strings.Join([]string{"ListAntitamperRuleResponse", string(data)}, " ")
}
