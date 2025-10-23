package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReportSummaryDataResponse Response Object
type ShowReportSummaryDataResponse struct {

	// 数据项名称，例如task-status-pie
	DataItemName *string `json:"data_item_name,omitempty"`

	// 数据点个数，例如10
	DataLength *int32 `json:"data_length,omitempty"`

	// 数据维度名称，例如区域、日期、任务类型、资源类型
	DimensionName *string `json:"dimension_name,omitempty"`

	// 数据维度取值，长度等于data_length，例如[\"success\",\"failed\",\"skipped\"]，[\"202401\",\"202402\",\"202403']
	DimensionData *[]string `json:"dimension_data,omitempty"`

	// 数据取值名称，例如[任务数量、任务成功率，资源数量、资源保护率、资源合规率
	ValueName *[]string `json:"value_name,omitempty"`

	// 数据取值，和value_name相对应，可以是多组数据，例如[[100,0.98,0.97],[99,0.98.0.99]]
	ValueData      *[][]float64 `json:"value_data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowReportSummaryDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReportSummaryDataResponse struct{}"
	}

	return strings.Join([]string{"ShowReportSummaryDataResponse", string(data)}, " ")
}
