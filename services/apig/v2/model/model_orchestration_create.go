package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OrchestrationCreate struct {

	// 编排映射规则名称。  支持英文，数字，下划线，且只能以英文开头，3-64个字符，同一实例下不允许重名。
	OrchestrationName string `json:"orchestration_name"`

	// 编排策略，有以下几种策略类型： - list：列表； - hash：哈希； - range：区间； - hash_range: 哈希+区间； - none_value: 空值映射； - default: 默认值映射； - head_n: 截取前n项； - tail_n: 截取后n项； 当编排策略为list时，orchestration_map列表长度*map_param_list长度不超过3000。
	OrchestrationStrategy OrchestrationCreateOrchestrationStrategy `json:"orchestration_strategy"`

	OrchestrationMappedParam *OrchestrationMappedParam `json:"orchestration_mapped_param,omitempty"`

	// 是否为预处理策略，预处理策略只会生成临时参数作为后面参数编排规则的入参标记。当为预处理策略时，该编排规则不能作为除default之外的最后一个编排规则。
	IsPreprocessing *bool `json:"is_preprocessing,omitempty"`

	// 编排映射规则列表，列表长度范围为1-300。 编排映射规则的生效优先级与列表顺序保持一致，列表中靠前的配置匹配优先级较高。 映射规则不能重复，当orchestration_strategy=list时，map_param_list的列表也不能包含重复元素。
	OrchestrationMap *[]OrchestrationMap `json:"orchestration_map,omitempty"`
}

func (o OrchestrationCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrchestrationCreate struct{}"
	}

	return strings.Join([]string{"OrchestrationCreate", string(data)}, " ")
}

type OrchestrationCreateOrchestrationStrategy struct {
	value string
}

type OrchestrationCreateOrchestrationStrategyEnum struct {
	LIST       OrchestrationCreateOrchestrationStrategy
	HASH       OrchestrationCreateOrchestrationStrategy
	RANGE      OrchestrationCreateOrchestrationStrategy
	HASH_RANGE OrchestrationCreateOrchestrationStrategy
	NONE_VALUE OrchestrationCreateOrchestrationStrategy
	DEFAULT    OrchestrationCreateOrchestrationStrategy
	HEAD_N     OrchestrationCreateOrchestrationStrategy
	TAIL_N     OrchestrationCreateOrchestrationStrategy
}

func GetOrchestrationCreateOrchestrationStrategyEnum() OrchestrationCreateOrchestrationStrategyEnum {
	return OrchestrationCreateOrchestrationStrategyEnum{
		LIST: OrchestrationCreateOrchestrationStrategy{
			value: "list",
		},
		HASH: OrchestrationCreateOrchestrationStrategy{
			value: "hash",
		},
		RANGE: OrchestrationCreateOrchestrationStrategy{
			value: "range",
		},
		HASH_RANGE: OrchestrationCreateOrchestrationStrategy{
			value: "hash_range",
		},
		NONE_VALUE: OrchestrationCreateOrchestrationStrategy{
			value: "none_value",
		},
		DEFAULT: OrchestrationCreateOrchestrationStrategy{
			value: "default",
		},
		HEAD_N: OrchestrationCreateOrchestrationStrategy{
			value: "head_n",
		},
		TAIL_N: OrchestrationCreateOrchestrationStrategy{
			value: "tail_n",
		},
	}
}

func (c OrchestrationCreateOrchestrationStrategy) Value() string {
	return c.value
}

func (c OrchestrationCreateOrchestrationStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrchestrationCreateOrchestrationStrategy) UnmarshalJSON(b []byte) error {
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
