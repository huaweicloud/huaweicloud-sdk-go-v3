package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobAlgorithmResponsePoliciesAutoSearchSearchParams struct {

	// 超参名称。
	Name *string `json:"name,omitempty"`

	// 参数类型。 - continuous：指定时表示这个超参是连续类型的。连续类型的超参在算法使用于训练作业时，控制台显示为输入框。 - discrete：指定时表示这个超参是离散类型的。离散类型的超参在算法使用于训练作业时，控制台显示为下拉选择框架。
	ParamType *string `json:"param_type,omitempty"`

	// 超参下界。
	LowerBound *string `json:"lower_bound,omitempty"`

	// 超参上界。
	UpperBound *string `json:"upper_bound,omitempty"`

	// 连续型超参离散化取值个数。
	DiscretePointsNum *string `json:"discrete_points_num,omitempty"`

	// 离散型超参的取值列表。
	DiscreteValues *[]string `json:"discrete_values,omitempty"`
}

func (o JobAlgorithmResponsePoliciesAutoSearchSearchParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobAlgorithmResponsePoliciesAutoSearchSearchParams struct{}"
	}

	return strings.Join([]string{"JobAlgorithmResponsePoliciesAutoSearchSearchParams", string(data)}, " ")
}
