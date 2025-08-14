package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectResourcesPolicyStatesSummaryResponse Response Object
type CollectResourcesPolicyStatesSummaryResponse struct {

	// 合规结果查询返回值
	Value *[]PolicyResourceComplianceSummary `json:"value,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CollectResourcesPolicyStatesSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectResourcesPolicyStatesSummaryResponse struct{}"
	}

	return strings.Join([]string{"CollectResourcesPolicyStatesSummaryResponse", string(data)}, " ")
}
