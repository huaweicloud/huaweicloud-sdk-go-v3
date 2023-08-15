package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WeakConstraintDto 分子约束
type WeakConstraintDto struct {

	// 属性名称
	Name *string `json:"name,omitempty"`

	Type *WeakConstraintType `json:"type"`

	// 属性约束类型bool的参数
	Bool *bool `json:"bool,omitempty"`

	// 属性约束类型range的参数
	Range *[]float32 `json:"range,omitempty"`

	Struct *StructureConstraintParamsDto `json:"struct,omitempty"`

	// 属性约束类型minimize和maximize的参数
	Quantiles *[]float32 `json:"quantiles,omitempty"`
}

func (o WeakConstraintDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WeakConstraintDto struct{}"
	}

	return strings.Join([]string{"WeakConstraintDto", string(data)}, " ")
}
