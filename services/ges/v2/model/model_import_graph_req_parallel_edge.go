package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportGraphReqParallelEdge 重复边处理。  图规格为（一千亿边）的图暂不支持该参数。
type ImportGraphReqParallelEdge struct {

	// 处理方式，取值为allow，ignore和override，默认为allow。  - allow表示允许重复边。 - ignore表示忽略之后的重复边。 - override表示覆盖之前的重复边。 图规格为（一千亿边）的图暂不支持该参数。
	Action *string `json:"action,omitempty"`

	// 重复边的定义，是否忽略Label。取值为true或者false，默认取true。  - true 表示重复边定义不包含Label，即用<源点，终点>标记一条边，不包含Label。 - false 表示重复边定义包含Label，即用<源点，终点，Label>标记一条边。 图规格为（一千亿边）的图暂不支持该参数。
	IgnoreLabel *bool `json:"ignore_label,omitempty"`
}

func (o ImportGraphReqParallelEdge) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportGraphReqParallelEdge struct{}"
	}

	return strings.Join([]string{"ImportGraphReqParallelEdge", string(data)}, " ")
}
