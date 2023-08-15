package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectConformancePackComplianceSummaryRequest Request Object
type CollectConformancePackComplianceSummaryRequest struct {

	// 最大的返回数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，通过上一个请求中返回的marker信息作为输入，获取当前页
	Marker *string `json:"marker,omitempty"`

	// 合规规则包名称。
	ConformancePackName *string `json:"conformance_pack_name,omitempty"`
}

func (o CollectConformancePackComplianceSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectConformancePackComplianceSummaryRequest struct{}"
	}

	return strings.Join([]string{"CollectConformancePackComplianceSummaryRequest", string(data)}, " ")
}
