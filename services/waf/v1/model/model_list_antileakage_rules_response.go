package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAntileakageRulesResponse Response Object
type ListAntileakageRulesResponse struct {

	// 防泄漏规则数量
	Total *int32 `json:"total,omitempty"`

	// 防泄漏规则列表
	Items          *[]LeakageListInfo `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAntileakageRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAntileakageRulesResponse struct{}"
	}

	return strings.Join([]string{"ListAntileakageRulesResponse", string(data)}, " ")
}
