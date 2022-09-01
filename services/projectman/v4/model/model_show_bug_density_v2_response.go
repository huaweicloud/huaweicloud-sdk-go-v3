package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBugDensityV2Response struct {

	// 指标分子数值
	DividendValue *string `json:"dividend_value,omitempty" xml:"dividend_value"`

	// 指标分母数值
	DivisorValue *string `json:"divisor_value,omitempty" xml:"divisor_value"`

	// 指标名称
	MetricName *string `json:"metric_name,omitempty" xml:"metric_name"`

	// 指标数值
	MetricValue *string `json:"metric_value,omitempty" xml:"metric_value"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 项目名称
	ProjectName    *string `json:"project_name,omitempty" xml:"project_name"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBugDensityV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBugDensityV2Response struct{}"
	}

	return strings.Join([]string{"ShowBugDensityV2Response", string(data)}, " ")
}
