package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 输入输出过滤配置
type StateDataFilter struct {

	// 输入过滤表达式（JsonPath）
	Input *string `json:"input,omitempty" xml:"input"`

	// 输出过滤表达式（JsonPath）
	Output *string `json:"output,omitempty" xml:"output"`
}

func (o StateDataFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StateDataFilter struct{}"
	}

	return strings.Join([]string{"StateDataFilter", string(data)}, " ")
}
