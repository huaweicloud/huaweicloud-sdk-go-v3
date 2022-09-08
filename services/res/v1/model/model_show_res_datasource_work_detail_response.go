package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowResDatasourceWorkDetailResponse struct {
	ErrorCount *ErrorCount `json:"error_count,omitempty"`

	DataStruct *DataStruct `json:"data_struct,omitempty"`

	// 错误样例(请求类型为DATA_INSPECTION时返回)。
	ErrorSamples *[]ErrorSample `json:"error_samples,omitempty"`

	// 数据检测结果(请求类型为DATA_INSPECTION时返回)。
	InspectResult *[]InspectResult `json:"inspect_result,omitempty"`

	// 是否成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 返回消息。
	Message *string `json:"message,omitempty"`

	// 合法率(请求类型为DATA_INSPECTION时返回)。
	LegalRate *float64 `json:"legal_rate,omitempty"`

	// 检测结果生成时间(请求类型为DATA_INSPECTION时返回)。
	InspectRstGeneratedTime *string `json:"inspect_rst_generated_time,omitempty"`

	FinalReport    *FinalReport `json:"final_report,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowResDatasourceWorkDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResDatasourceWorkDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowResDatasourceWorkDetailResponse", string(data)}, " ")
}
