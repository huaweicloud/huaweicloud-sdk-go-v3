package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 脱敏策略列表，每个策略对应一个字段，脱敏策略数最多100个。详情见“动态脱敏策略配置”。
type MaskStrategies struct {

	// 需要脱敏的字段名称，最大支持长度256。
	Name string `json:"name"`

	// 脱敏算法名称，详情见附录\"动态脱敏策略配置\"。
	Algorithm MaskStrategiesAlgorithm `json:"algorithm"`

	// 脱敏算法参数，详情见附录\"动态脱敏策略配置\"。
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

func (o MaskStrategies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MaskStrategies struct{}"
	}

	return strings.Join([]string{"MaskStrategies", string(data)}, " ")
}

type MaskStrategiesAlgorithm struct {
	value string
}

type MaskStrategiesAlgorithmEnum struct {
	SHA256  MaskStrategiesAlgorithm
	SHA512  MaskStrategiesAlgorithm
	AES     MaskStrategiesAlgorithm
	PRESNM  MaskStrategiesAlgorithm
	MASKNM  MaskStrategiesAlgorithm
	PRESXY  MaskStrategiesAlgorithm
	MASKXY  MaskStrategiesAlgorithm
	SYMBOL  MaskStrategiesAlgorithm
	KEYWORD MaskStrategiesAlgorithm
	NUMERIC MaskStrategiesAlgorithm
}

func GetMaskStrategiesAlgorithmEnum() MaskStrategiesAlgorithmEnum {
	return MaskStrategiesAlgorithmEnum{
		SHA256: MaskStrategiesAlgorithm{
			value: "SHA256",
		},
		SHA512: MaskStrategiesAlgorithm{
			value: "SHA512",
		},
		AES: MaskStrategiesAlgorithm{
			value: "AES",
		},
		PRESNM: MaskStrategiesAlgorithm{
			value: "PRESNM",
		},
		MASKNM: MaskStrategiesAlgorithm{
			value: "MASKNM",
		},
		PRESXY: MaskStrategiesAlgorithm{
			value: "PRESXY",
		},
		MASKXY: MaskStrategiesAlgorithm{
			value: "MASKXY",
		},
		SYMBOL: MaskStrategiesAlgorithm{
			value: "SYMBOL",
		},
		KEYWORD: MaskStrategiesAlgorithm{
			value: "KEYWORD",
		},
		NUMERIC: MaskStrategiesAlgorithm{
			value: "NUMERIC",
		},
	}
}

func (c MaskStrategiesAlgorithm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MaskStrategiesAlgorithm) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
