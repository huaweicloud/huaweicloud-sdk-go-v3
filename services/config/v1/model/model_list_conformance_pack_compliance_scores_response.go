package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConformancePackComplianceScoresResponse Response Object
type ListConformancePackComplianceScoresResponse struct {

	// 合规规则包分数查询列表。
	Value *[]ConformancePackScore `json:"value,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListConformancePackComplianceScoresResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConformancePackComplianceScoresResponse struct{}"
	}

	return strings.Join([]string{"ListConformancePackComplianceScoresResponse", string(data)}, " ")
}
