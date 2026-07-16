package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoSearchAlgoConfigParameter 训练作业、算法依赖参数。
type AutoSearchAlgoConfigParameter struct {

	// 参数键。
	Key *string `json:"key,omitempty"`

	// 参数值。
	Value *string `json:"value,omitempty"`

	// 参数种类。
	Type *string `json:"type,omitempty"`
}

func (o AutoSearchAlgoConfigParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoSearchAlgoConfigParameter struct{}"
	}

	return strings.Join([]string{"AutoSearchAlgoConfigParameter", string(data)}, " ")
}
