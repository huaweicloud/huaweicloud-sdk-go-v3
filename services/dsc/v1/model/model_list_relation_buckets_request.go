package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRelationBucketsRequest Request Object
type ListRelationBucketsRequest struct {

	// 任务ID
	JobId string `json:"job_id"`

	// 资产名称
	AssetsName *string `json:"assets_name,omitempty"`

	// 起始风险等级
	RiskStart int32 `json:"risk_start"`

	// 终止风险等级
	RiskEnd int32 `json:"risk_end"`

	// 页码
	Offset *int32 `json:"offset,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRelationBucketsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelationBucketsRequest struct{}"
	}

	return strings.Join([]string{"ListRelationBucketsRequest", string(data)}, " ")
}
