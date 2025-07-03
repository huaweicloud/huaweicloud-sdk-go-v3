package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlowDetailRespBody 响应流详情
type FlowDetailRespBody struct {

	// 流名称
	Name string `json:"name"`

	// 区域
	Region string `json:"region"`

	// 入流信息
	Sources []FlowSource `json:"sources"`

	// 状态
	Status string `json:"status"`

	// 流id
	FlowId string `json:"flow_id"`

	// 流输出信息
	Outputs *[]FlowOutput `json:"outputs,omitempty"`
}

func (o FlowDetailRespBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowDetailRespBody struct{}"
	}

	return strings.Join([]string{"FlowDetailRespBody", string(data)}, " ")
}
