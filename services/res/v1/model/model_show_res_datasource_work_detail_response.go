package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowResDatasourceWorkDetailResponse struct {
	ErrorCount *ErrorCount `json:"error_count,omitempty" xml:"error_count"`

	DataStruct *DataStruct `json:"data_struct,omitempty" xml:"data_struct"`

	// 错误样例(请求类型为DATA_INSPECTION时返回)。
	ErrorSamples *[]ErrorSample `json:"error_samples,omitempty" xml:"error_samples"`

	// 数据检测结果(请求类型为DATA_INSPECTION时返回)。
	InspectResult *[]InspectResult `json:"inspect_result,omitempty" xml:"inspect_result"`

	// 是否成功。
	IsSuccess *bool `json:"is_success,omitempty" xml:"is_success"`

	// 返回消息。
	Message *string `json:"message,omitempty" xml:"message"`

	// 合法率(请求类型为DATA_INSPECTION时返回)。
	LegalRate *float64 `json:"legal_rate,omitempty" xml:"legal_rate"`

	// 检测结果生成时间(请求类型为DATA_INSPECTION时返回)。
	InspectRstGeneratedTime *string `json:"inspect_rst_generated_time,omitempty" xml:"inspect_rst_generated_time"`

	FinalReport    *FinalReport `json:"final_report,omitempty" xml:"final_report"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowResDatasourceWorkDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResDatasourceWorkDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowResDatasourceWorkDetailResponse", string(data)}, " ")
}
