package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBugsPerDeveloperResponse struct {

	// 指标分子数值
	DividendValue *string `json:"dividend_value,omitempty"`

	// 指标分母数值
	DivisorValue *string `json:"divisor_value,omitempty"`

	// 指标名称
	MetricName *string `json:"metric_name,omitempty"`

	// 指标数值
	MetricValue *string `json:"metric_value,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 项目名称
	ProjectName    *string `json:"project_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBugsPerDeveloperResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBugsPerDeveloperResponse struct{}"
	}

	return strings.Join([]string{"ShowBugsPerDeveloperResponse", string(data)}, " ")
}
