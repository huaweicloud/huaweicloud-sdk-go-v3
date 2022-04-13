package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 诊断结论
type ConclusionItem struct {
	// 结论id

	Id int32 `json:"id"`
	// 结论参数

	Params map[string]string `json:"params,omitempty"`
}

func (o ConclusionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConclusionItem struct{}"
	}

	return strings.Join([]string{"ConclusionItem", string(data)}, " ")
}
