package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StrongConstraintDto 分子约束
type StrongConstraintDto struct {

	// 属性名称
	Name *string `json:"name,omitempty"`

	Type *StrongConstraintType `json:"type"`

	// 属性约束类型bool的参数
	Bool *bool `json:"bool,omitempty"`

	// 属性约束类型range的参数
	Range *[]float32 `json:"range,omitempty"`

	Struct *StructureConstraintParamsDto `json:"struct,omitempty"`
}

func (o StrongConstraintDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StrongConstraintDto struct{}"
	}

	return strings.Join([]string{"StrongConstraintDto", string(data)}, " ")
}
