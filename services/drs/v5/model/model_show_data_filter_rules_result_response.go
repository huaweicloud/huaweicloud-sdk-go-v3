package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataFilterRulesResultResponse Response Object
type ShowDataFilterRulesResultResponse struct {

	// 总数。
	TotalCount *int64 `json:"total_count,omitempty"`

	// 数据过滤规则列表。
	Rules *[]DataFilterRuleResp `json:"rules,omitempty"`

	// 查询过滤规则进度。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDataFilterRulesResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataFilterRulesResultResponse struct{}"
	}

	return strings.Join([]string{"ShowDataFilterRulesResultResponse", string(data)}, " ")
}
