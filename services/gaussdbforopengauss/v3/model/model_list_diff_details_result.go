package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 参数配置模板差异列表
type ListDiffDetailsResult struct {

	// 参数名称。
	Name *string `json:"name,omitempty"`

	// 比较参数组的参数值。
	SourceValue *string `json:"source_value,omitempty"`

	// 目标参数组的参数值。
	TargetValue *string `json:"target_value,omitempty"`
}

func (o ListDiffDetailsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiffDetailsResult struct{}"
	}

	return strings.Join([]string{"ListDiffDetailsResult", string(data)}, " ")
}
