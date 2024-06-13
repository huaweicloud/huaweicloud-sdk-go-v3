package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Input struct {

	// 参数名
	Name string `json:"name"`

	// 参数值
	Value *interface{} `json:"value"`

	// 值类型
	Type *string `json:"type,omitempty"`
}

func (o Input) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Input struct{}"
	}

	return strings.Join([]string{"Input", string(data)}, " ")
}
