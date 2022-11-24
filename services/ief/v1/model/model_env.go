package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 环境变量
type Env struct {

	// 环境变量的key，由大小写字母或下划线开头，由数字、大小写字母、下划线组成，最大长度2048个字符，不允许重复
	Name string `json:"name"`

	// 环境变量的value，最大长度20480个字符。value、value_from和field_path必须三选一使用。
	Value *string `json:"value,omitempty"`

	ValueFrom *ValueFrom `json:"value_from,omitempty"`

	// 该参数目前只支持赋值\"status.hostIP\"，即引用边缘节点的IP地址作为环境变量。
	FieldPath *string `json:"field_path,omitempty"`
}

func (o Env) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Env struct{}"
	}

	return strings.Join([]string{"Env", string(data)}, " ")
}
