package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConditionInstance struct {

	// 条件索引，用于标识当前处理的是哪个条件
	ConditionIndex *int32 `json:"conditionIndex,omitempty"`

	// 条件映射，键为整数，值为字符串列表，用于存储各个条件的信息
	ConditionMap map[string][]string `json:"conditionMap,omitempty"`

	// 依赖参数实例的映射
	DepParamMap map[string]DepParamInstance `json:"depParamMap,omitempty"`

	// 声明索引，用于标识当前处理的是哪个声明
	StatementIndex *int32 `json:"statementIndex,omitempty"`
}

func (o ConditionInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConditionInstance struct{}"
	}

	return strings.Join([]string{"ConditionInstance", string(data)}, " ")
}
