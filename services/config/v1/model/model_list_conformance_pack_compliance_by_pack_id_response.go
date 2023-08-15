package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConformancePackComplianceByPackIdResponse Response Object
type ListConformancePackComplianceByPackIdResponse struct {

	// 合规规则包的合规规则评估结果列表。
	Value *[]ConformancePackCompliance `json:"value,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListConformancePackComplianceByPackIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConformancePackComplianceByPackIdResponse struct{}"
	}

	return strings.Join([]string{"ListConformancePackComplianceByPackIdResponse", string(data)}, " ")
}
