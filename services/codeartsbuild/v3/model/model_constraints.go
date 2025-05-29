package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Constraints 约束
type Constraints struct {

	// 错误信息
	Errormsg *string `json:"errormsg,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 参考值
	Value *string `json:"value,omitempty"`
}

func (o Constraints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Constraints struct{}"
	}

	return strings.Join([]string{"Constraints", string(data)}, " ")
}
