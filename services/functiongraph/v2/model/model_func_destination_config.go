package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 函数通知目标参数配置。
type FuncDestinationConfig struct {

	// 目标类型。  - OBS：通知到OBS服务。 - SMN：通知到SMN服务。 - DIS：通知到DIS服务。 - FunctionGraph： 通知到函数服务。
	Destination *FuncDestinationConfigDestination `json:"destination,omitempty"`

	// 通知目标服务对应参数,json字符串。  - OBS：包含bucket桶，对象目录前缀prefix，对象默认expires过期时间[0~365]天，0默认不过期。 - SMN：包含smn 主题topic_urn。 - DIS：包含DIS 通道名stream_name。 - FunctionGraph：包含func_urn，函数urn
	Param *string `json:"param,omitempty"`
}

func (o FuncDestinationConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FuncDestinationConfig struct{}"
	}

	return strings.Join([]string{"FuncDestinationConfig", string(data)}, " ")
}

type FuncDestinationConfigDestination struct {
	value string
}

type FuncDestinationConfigDestinationEnum struct {
	OBS            FuncDestinationConfigDestination
	SMN            FuncDestinationConfigDestination
	DIS            FuncDestinationConfigDestination
	FUNCTION_GRAPH FuncDestinationConfigDestination
}

func GetFuncDestinationConfigDestinationEnum() FuncDestinationConfigDestinationEnum {
	return FuncDestinationConfigDestinationEnum{
		OBS: FuncDestinationConfigDestination{
			value: "OBS",
		},
		SMN: FuncDestinationConfigDestination{
			value: "SMN",
		},
		DIS: FuncDestinationConfigDestination{
			value: "DIS",
		},
		FUNCTION_GRAPH: FuncDestinationConfigDestination{
			value: "FunctionGraph",
		},
	}
}

func (c FuncDestinationConfigDestination) Value() string {
	return c.value
}

func (c FuncDestinationConfigDestination) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FuncDestinationConfigDestination) UnmarshalJSON(b []byte) error {
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
