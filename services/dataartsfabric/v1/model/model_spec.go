package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Spec 规格信息
type Spec struct {

	// 资源规格，从规格列表查询获取。
	SpecCode *string `json:"spec_code,omitempty"`

	SpecType *SpecType `json:"spec_type,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 步长
	Stride *int32 `json:"stride,omitempty"`

	// 单位
	Unit *string `json:"unit,omitempty"`

	// 资源步数，最小为10000，步长为1。输入范围还需要满足规格列表接口的步长白名单
	MinStrideNum *int32 `json:"min_stride_num,omitempty"`

	// 资源步数，最小为10000，步长为1。输入范围还需要满足规格列表接口的步长白名单
	MaxStrideNum *int32 `json:"max_stride_num,omitempty"`

	// 使用量单位标识
	UsageMeasureId *int32 `json:"usage_measure_id,omitempty"`

	// 使用量因子
	UsageFactor *string `json:"usage_factor,omitempty"`

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
