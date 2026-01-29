package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAdhocResultResponse Response Object
type ShowAdhocResultResponse struct {

	// 获取数据的批次，为0则为第一次查询
	Batch *int32 `json:"batch,omitempty"`

	// 统计分析结果字段类型
	Schema *[]AdhocQueryAnalysisField `json:"schema,omitempty"`

	// 统计分析结果数据
	Datarows *[][]interface{} `json:"datarows,omitempty"`

	// 统计分析结果数据
	DatarowsUpsert *[][]DataRow `json:"datarows_upsert,omitempty"`

	// 统计分析结果总数
	Total *int32 `json:"total,omitempty"`

	// 返回的统计分析结果条数
	Size *int32 `json:"size,omitempty"`

	// 是否有下一批数据
	Next *int32 `json:"next,omitempty"`

	Tips *ShowAdhocQueryResultResponseBodyTips `json:"tips,omitempty"`

	// UUID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAdhocResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAdhocResultResponse struct{}"
	}

	return strings.Join([]string{"ShowAdhocResultResponse", string(data)}, " ")
}
