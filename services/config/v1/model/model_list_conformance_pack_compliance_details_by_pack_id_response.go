package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConformancePackComplianceDetailsByPackIdResponse Response Object
type ListConformancePackComplianceDetailsByPackIdResponse struct {

	// 合规规则包的合规规则评估结果详情列表。
	Value *[]ConformancePackComplianceDetail `json:"value,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListConformancePackComplianceDetailsByPackIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConformancePackComplianceDetailsByPackIdResponse struct{}"
	}

	return strings.Join([]string{"ListConformancePackComplianceDetailsByPackIdResponse", string(data)}, " ")
}
