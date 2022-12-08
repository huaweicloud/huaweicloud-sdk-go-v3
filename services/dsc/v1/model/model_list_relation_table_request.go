package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRelationTableRequest struct {

	// 任务ID
	JobId string `json:"job_id"`

	// 数据库ID
	DbId string `json:"db_id"`

	// 资产名称
	AssetsName *string `json:"assets_name,omitempty"`

	// 起始风险等级
	RiskStart int32 `json:"risk_start"`

	// 终止风险等级
	RiskEnd int32 `json:"risk_end"`

	// 偏移量
	Offset int32 `json:"offset"`

	// 页面大小
	Size int32 `json:"size"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRelationTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelationTableRequest struct{}"
	}

	return strings.Join([]string{"ListRelationTableRequest", string(data)}, " ")
}
