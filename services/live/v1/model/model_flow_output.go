package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlowOutput 流输出信息
type FlowOutput struct {

	// 名称
	Name string `json:"name"`

	// 类型
	Type string `json:"type"`
}

func (o FlowOutput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowOutput struct{}"
	}

	return strings.Join([]string{"FlowOutput", string(data)}, " ")
}
