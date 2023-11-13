package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Item struct {

	// 诊断结果
	Result *string `json:"result,omitempty"`

	// 诊断报告ID
	Id *string `json:"id,omitempty"`
}

func (o Item) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Item struct{}"
	}

	return strings.Join([]string{"Item", string(data)}, " ")
}
