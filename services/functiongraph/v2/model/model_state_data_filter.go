package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StateDataFilter 输入输出过滤配置
type StateDataFilter struct {

	// 输入过滤表达式（JsonPath）
	Input *string `json:"input,omitempty"`

	// 输出过滤表达式（JsonPath）
	Output *string `json:"output,omitempty"`
}

func (o StateDataFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StateDataFilter struct{}"
	}

	return strings.Join([]string{"StateDataFilter", string(data)}, " ")
}
