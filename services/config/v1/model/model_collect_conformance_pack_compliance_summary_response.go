package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectConformancePackComplianceSummaryResponse Response Object
type CollectConformancePackComplianceSummaryResponse struct {

	// 合规规则包的合规结果概览列表。
	Value *[]ConformancePackComplianceSummary `json:"value,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CollectConformancePackComplianceSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectConformancePackComplianceSummaryResponse struct{}"
	}

	return strings.Join([]string{"CollectConformancePackComplianceSummaryResponse", string(data)}, " ")
}
