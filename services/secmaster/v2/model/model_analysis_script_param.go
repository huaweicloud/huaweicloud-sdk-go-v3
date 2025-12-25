package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AnalysisScriptParam struct {

	// 键 key
	Key *string `json:"key,omitempty"`

	// 值 value
	Value *string `json:"value,omitempty"`
}

func (o AnalysisScriptParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnalysisScriptParam struct{}"
	}

	return strings.Join([]string{"AnalysisScriptParam", string(data)}, " ")
}
