package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Spec 规格信息
type Spec struct {

	// 规格编码。lakeformation.unit.basic.qps：每秒查询率（QPS）产品
	SpecCode *string `json:"spec_code,omitempty"`

	// 资源编码
	ResourceType *string `json:"resource_type,omitempty"`

	// 步长
	Stride *int32 `json:"stride,omitempty"`

	// 单位
	Unit *string `json:"unit,omitempty"`

	// 最小步数
	MinStrideNum *int32 `json:"min_stride_num,omitempty"`

	// 最大步数
	MaxStrideNum *int32 `json:"max_stride_num,omitempty"`

	// 使用量单位标识
	UsageMeasureId *int32 `json:"usage_measure_id,omitempty"`

	// 使用量因子
	UsageFactor *string `json:"usage_factor,omitempty"`

	// 使用量，包含免费额度和单位额度，例如api调用次数，单位是次，前100万次调用免费，计费标准是5元每100万次，这里返回200万，元数据个数，单位是万个，前100万个免费，计费标准是5元每10万个，这里返回110
	UsageValue *int32 `json:"usage_value,omitempty"`

	// 免费使用额度，例如api调用次数，单位是次，前100万次调用免费，这里返回100万，元数据个数，单位是万个，前100万个免费，这里返回100
	FreeUsageValue *int32 `json:"free_usage_value,omitempty"`

	// 步数白名单，返回时，步数必须是白名单中的值
	StrideNumWhitelist *[]int32 `json:"stride_num_whitelist,omitempty"`
}

func (o Spec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Spec struct{}"
	}

	return strings.Join([]string{"Spec", string(data)}, " ")
}
