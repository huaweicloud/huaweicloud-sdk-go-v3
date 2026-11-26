package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DependArg **参数解释**：依赖的单个参数信息
type DependArg struct {

	// **参数解释**：依赖的参数key值
	Key string `json:"key"`

	// **参数解释**：依赖的参数value值,多个值用','分隔
	Value string `json:"value"`
}

func (o DependArg) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DependArg struct{}"
	}

	return strings.Join([]string{"DependArg", string(data)}, " ")
}
