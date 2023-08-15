package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRelationColumnRequest Request Object
type ListRelationColumnRequest struct {

	// 任务ID
	JobId string `json:"job_id"`

	// 表ID
	TableId string `json:"table_id"`

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

func (o ListRelationColumnRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelationColumnRequest struct{}"
	}

	return strings.Join([]string{"ListRelationColumnRequest", string(data)}, " ")
}
