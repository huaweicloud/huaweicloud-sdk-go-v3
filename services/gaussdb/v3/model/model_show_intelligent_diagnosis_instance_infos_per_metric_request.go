package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIntelligentDiagnosisInstanceInfosPerMetricRequest Request Object
type ShowIntelligentDiagnosisInstanceInfosPerMetricRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 指标名。
	MetricName string `json:"metric_name"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset int32 `json:"offset"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit int32 `json:"limit"`
}

func (o ShowIntelligentDiagnosisInstanceInfosPerMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIntelligentDiagnosisInstanceInfosPerMetricRequest struct{}"
	}

	return strings.Join([]string{"ShowIntelligentDiagnosisInstanceInfosPerMetricRequest", string(data)}, " ")
}
