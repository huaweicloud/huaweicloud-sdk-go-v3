package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAntitamperRuleResponse Response Object
type ListAntitamperRuleResponse struct {

	// 网页防篡改规则总条数
	Total *int32 `json:"total,omitempty"`

	// 网页防篡改规则数组
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
