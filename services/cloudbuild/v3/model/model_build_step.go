package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BuildStep struct {

	// 步骤名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 步骤状态,可选值（running运行中，success成功，error失败，未运行为空字符串）
	Status *string `json:"status,omitempty" xml:"status"`

	// 步骤执行时长，单位ms
	BuildTime *int32 `json:"build_time,omitempty" xml:"build_time"`
}

func (o BuildStep) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildStep struct{}"
	}

	return strings.Join([]string{"BuildStep", string(data)}, " ")
}
