package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ParallelEdge struct {

	// 处理方式，取值为allow，ignore和override，默认为allow。 - allow表示允许重复边。 - ignore表示忽略之后的重复边。 - override表示覆盖之前的重复边。
	Action *string `json:"action,omitempty"`

	// 重复边的定义，是否忽略Label。取值为true或者false，默认取true。 - true 表示重复边定义不包含Label，即用<源点，终点>标记一条边，不包含Label。 - false 表示重复边定义包含Label，即用<源点，终点，Label>标记一条边。
	IgnoreLabel *bool `json:"ignoreLabel,omitempty"`
}

func (o ParallelEdge) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParallelEdge struct{}"
	}

	return strings.Join([]string{"ParallelEdge", string(data)}, " ")
}
