package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OrchestrationBaseInfo 编排规则。  单个实例允许创建的编排规则有配额限制，具体请参见产品介绍的“配额说明”章节。
type OrchestrationBaseInfo struct {

	// 编排映射规则名称。  支持英文，数字，下划线，且只能以英文开头，3-64个字符，同一实例下不允许重名。
	OrchestrationName string `json:"orchestration_name"`

	// 编排策略，有以下几种策略类型： - list：列表； - hash：哈希； - range：区间； - hash_range: 哈希+区间； - none_value: 空值映射； - default: 默认值映射； - head_n: 截取前n项； - tail_n: 截取后n项； 当编排策略为list时，orchestration_map列表长度*map_param_list长度不超过3000。
	OrchestrationStrategy OrchestrationBaseInfoOrchestrationStrategy `json:"orchestration_strategy"`

	OrchestrationMappedParam *OrchestrationMappedParam `json:"orchestration_mapped_param,omitempty"`

	// 是否为预处理策略，预处理策略只会生成临时参数作为后面参数编排规则的入参标记。当为预处理策略时，该编排规则不能作为除default之外的最后一个编排规则。
	IsPreprocessing *bool `json:"is_preprocessing,omitempty"`
}

func (o OrchestrationBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrchestrationBaseInfo struct{}"
	}

	return strings.Join([]string{"OrchestrationBaseInfo", string(data)}, " ")
}

type OrchestrationBaseInfoOrchestrationStrategy struct {
	value string
}

type OrchestrationBaseInfoOrchestrationStrategyEnum struct {
	LIST       OrchestrationBaseInfoOrchestrationStrategy
	HASH       OrchestrationBaseInfoOrchestrationStrategy
	RANGE      OrchestrationBaseInfoOrchestrationStrategy
	HASH_RANGE OrchestrationBaseInfoOrchestrationStrategy
	NONE_VALUE OrchestrationBaseInfoOrchestrationStrategy
	DEFAULT    OrchestrationBaseInfoOrchestrationStrategy
	HEAD_N     OrchestrationBaseInfoOrchestrationStrategy
	TAIL_N     OrchestrationBaseInfoOrchestrationStrategy
}

func GetOrchestrationBaseInfoOrchestrationStrategyEnum() OrchestrationBaseInfoOrchestrationStrategyEnum {
	return OrchestrationBaseInfoOrchestrationStrategyEnum{
		LIST: OrchestrationBaseInfoOrchestrationStrategy{
			value: "list",
		},
		HASH: OrchestrationBaseInfoOrchestrationStrategy{
			value: "hash",
		},
		RANGE: OrchestrationBaseInfoOrchestrationStrategy{
			value: "range",
		},
		HASH_RANGE: OrchestrationBaseInfoOrchestrationStrategy{
			value: "hash_range",
		},
		NONE_VALUE: OrchestrationBaseInfoOrchestrationStrategy{
			value: "none_value",
		},
		DEFAULT: OrchestrationBaseInfoOrchestrationStrategy{
			value: "default",
		},
		HEAD_N: OrchestrationBaseInfoOrchestrationStrategy{
			value: "head_n",
		},
		TAIL_N: OrchestrationBaseInfoOrchestrationStrategy{
			value: "tail_n",
		},
	}
}

func (c OrchestrationBaseInfoOrchestrationStrategy) Value() string {
	return c.value
}

func (c OrchestrationBaseInfoOrchestrationStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrchestrationBaseInfoOrchestrationStrategy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
