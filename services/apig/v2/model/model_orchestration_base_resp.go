package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// OrchestrationBaseResp 编排规则基本信息详情。
type OrchestrationBaseResp struct {

	// 编排映射规则名称。  支持英文，数字，下划线，且只能以英文开头，3-64个字符，同一实例下不允许重名。
	OrchestrationName string `json:"orchestration_name"`

	// 编排策略，有以下几种策略类型： - list：列表； - hash：哈希； - range：区间； - hash_range: 哈希+区间； - none_value: 空值映射； - default: 默认值映射； - head_n: 截取前n项； - tail_n: 截取后n项； 当编排策略为list时，orchestration_map列表长度*map_param_list长度不超过3000。
	OrchestrationStrategy OrchestrationBaseRespOrchestrationStrategy `json:"orchestration_strategy"`

	OrchestrationMappedParam *OrchestrationMappedParam `json:"orchestration_mapped_param,omitempty"`

	// 是否为预处理策略，预处理策略只会生成临时参数作为后面参数编排规则的入参标记。当为预处理策略时，该编排规则不能作为除default之外的最后一个编排规则。
	IsPreprocessing *bool `json:"is_preprocessing,omitempty"`

	// 编排规则编号。
	OrchestrationId *string `json:"orchestration_id,omitempty"`

	// 编排规则创建时间。
	OrchestrationCreateTime *sdktime.SdkTime `json:"orchestration_create_time,omitempty"`

	// 编排规则更新时间。
	OrchestrationUpdateTime *sdktime.SdkTime `json:"orchestration_update_time,omitempty"`
}

func (o OrchestrationBaseResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrchestrationBaseResp struct{}"
	}

	return strings.Join([]string{"OrchestrationBaseResp", string(data)}, " ")
}

type OrchestrationBaseRespOrchestrationStrategy struct {
	value string
}

type OrchestrationBaseRespOrchestrationStrategyEnum struct {
	LIST       OrchestrationBaseRespOrchestrationStrategy
	HASH       OrchestrationBaseRespOrchestrationStrategy
	RANGE      OrchestrationBaseRespOrchestrationStrategy
	HASH_RANGE OrchestrationBaseRespOrchestrationStrategy
	NONE_VALUE OrchestrationBaseRespOrchestrationStrategy
	DEFAULT    OrchestrationBaseRespOrchestrationStrategy
	HEAD_N     OrchestrationBaseRespOrchestrationStrategy
	TAIL_N     OrchestrationBaseRespOrchestrationStrategy
}

func GetOrchestrationBaseRespOrchestrationStrategyEnum() OrchestrationBaseRespOrchestrationStrategyEnum {
	return OrchestrationBaseRespOrchestrationStrategyEnum{
		LIST: OrchestrationBaseRespOrchestrationStrategy{
			value: "list",
		},
		HASH: OrchestrationBaseRespOrchestrationStrategy{
			value: "hash",
		},
		RANGE: OrchestrationBaseRespOrchestrationStrategy{
			value: "range",
		},
		HASH_RANGE: OrchestrationBaseRespOrchestrationStrategy{
			value: "hash_range",
		},
		NONE_VALUE: OrchestrationBaseRespOrchestrationStrategy{
			value: "none_value",
		},
		DEFAULT: OrchestrationBaseRespOrchestrationStrategy{
			value: "default",
		},
		HEAD_N: OrchestrationBaseRespOrchestrationStrategy{
			value: "head_n",
		},
		TAIL_N: OrchestrationBaseRespOrchestrationStrategy{
			value: "tail_n",
		},
	}
}

func (c OrchestrationBaseRespOrchestrationStrategy) Value() string {
	return c.value
}

func (c OrchestrationBaseRespOrchestrationStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OrchestrationBaseRespOrchestrationStrategy) UnmarshalJSON(b []byte) error {
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
