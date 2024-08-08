package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TargetOptLigand 配体设置
type TargetOptLigand struct {
	File *ProbeDrugFile `json:"file"`

	// 配体力场, 支持选择gaff, gaff2
	ForceField *string `json:"force_field,omitempty"`
}

func (o TargetOptLigand) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetOptLigand struct{}"
	}

	return strings.Join([]string{"TargetOptLigand", string(data)}, " ")
}
