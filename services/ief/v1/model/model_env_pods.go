package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnvPods 环境变量
type EnvPods struct {

	// 环境变量的key，由大小写字母或下划线开头，由数字、大小写字母、下划线组成，最大长度2048个字符，不允许重复
	Name string `json:"name"`

	// 环境变量的value，最大长度20480个字符。value、value_from和field_path必须三选一使用。
	Value *string `json:"value,omitempty"`
}

func (o EnvPods) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvPods struct{}"
	}

	return strings.Join([]string{"EnvPods", string(data)}, " ")
}
