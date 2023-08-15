package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StructureConstraintParamsDto 结构约束参数
type StructureConstraintParamsDto struct {

	// 子结构SMILES
	Structs []string `json:"structs"`

	// 是否排除子结构
	Exclusive bool `json:"exclusive"`

	Operator *OperatorType `json:"operator,omitempty"`
}

func (o StructureConstraintParamsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructureConstraintParamsDto struct{}"
	}

	return strings.Join([]string{"StructureConstraintParamsDto", string(data)}, " ")
}
