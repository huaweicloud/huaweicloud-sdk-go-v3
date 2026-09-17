package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRapidGrowthTablesResponse Response Object
type ListRapidGrowthTablesResponse struct {

	// 异常增长表信息列表
	Tables *[]RapidGrowthTableInfo `json:"tables,omitempty"`

	// 诊断阈值
	Threshold *int64 `json:"threshold,omitempty"`

	// 上次诊断时间
	LastDiagnoseTimestamp *int64 `json:"last_diagnose_timestamp,omitempty"`

	// 最近一次诊断时间
	First2LastTimestamp *int64 `json:"first2_last_timestamp,omitempty"`

	// 最近第二次诊断时间
	Second2LastTimestamp *int64 `json:"second2_last_timestamp,omitempty"`
	HttpStatusCode       int    `json:"-"`
}

func (o ListRapidGrowthTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRapidGrowthTablesResponse struct{}"
	}

	return strings.Join([]string{"ListRapidGrowthTablesResponse", string(data)}, " ")
}
