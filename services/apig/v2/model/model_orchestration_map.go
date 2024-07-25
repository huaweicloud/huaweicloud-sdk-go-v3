package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OrchestrationMap 编排映射规则。
type OrchestrationMap struct {

	// 用于映射编排后参数的列表配置，当orchestration_strategy=list时必填，列表长度范围为0-3000。  列表的取值只支持英文，数字，中划线和下划线，1-128个字符。
	MapParamList *[]string `json:"map_param_list,omitempty"`

	MapParamRange *OrchestrationMapParamRange `json:"map_param_range,omitempty"`

	// 编排后的参数取值，只支持英文和数字，1-128个字符。 当orchestration_strategy为hash、head_n、tail_n，或者is_preprocessing为false时，非必填，其他情况必填。
	MappedParamValue *string `json:"mapped_param_value,omitempty"`

	// 截取长度，取值范围为1-100，当策略类型为head_n和tail_n时必填，当截取长度大于参数长度时，截取参数的结果为完整参数。
	InterceptLength *int32 `json:"intercept_length,omitempty"`
}

func (o OrchestrationMap) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrchestrationMap struct{}"
	}

	return strings.Join([]string{"OrchestrationMap", string(data)}, " ")
}
