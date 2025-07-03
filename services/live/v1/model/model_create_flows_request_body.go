package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlowsRequestBody 创建流请求体
type CreateFlowsRequestBody struct {

	// 流名称
	Name string `json:"name"`

	// 区域
	Region string `json:"region"`

	// 入流信息
	Sources []FlowSource `json:"sources"`
}

func (o CreateFlowsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlowsRequestBody struct{}"
	}

	return strings.Join([]string{"CreateFlowsRequestBody", string(data)}, " ")
}
