package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BuildParamsBodyResult 结果
type BuildParamsBodyResult struct {

	// 构建参数约束列表
	BuildParameters *[]BuildParams `json:"build_parameters,omitempty"`
}

func (o BuildParamsBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildParamsBodyResult struct{}"
	}

	return strings.Join([]string{"BuildParamsBodyResult", string(data)}, " ")
}
