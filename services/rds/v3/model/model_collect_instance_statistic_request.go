package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CollectInstanceStatisticRequest Request Object
type CollectInstanceStatisticRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 引擎类型
	Engine *CollectInstanceStatisticRequestEngine `json:"engine,omitempty"`
}

func (o CollectInstanceStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectInstanceStatisticRequest struct{}"
	}

	return strings.Join([]string{"CollectInstanceStatisticRequest", string(data)}, " ")
}

type CollectInstanceStatisticRequestEngine struct {
	value string
}

type CollectInstanceStatisticRequestEngineEnum struct {
	MYSQL      CollectInstanceStatisticRequestEngine
	POSTGRESQL CollectInstanceStatisticRequestEngine
	SQLSERVER  CollectInstanceStatisticRequestEngine
}

func GetCollectInstanceStatisticRequestEngineEnum() CollectInstanceStatisticRequestEngineEnum {
	return CollectInstanceStatisticRequestEngineEnum{
		MYSQL: CollectInstanceStatisticRequestEngine{
			value: "mysql",
		},
		POSTGRESQL: CollectInstanceStatisticRequestEngine{
			value: "postgresql",
		},
		SQLSERVER: CollectInstanceStatisticRequestEngine{
			value: "sqlserver",
		},
	}
}

func (c CollectInstanceStatisticRequestEngine) Value() string {
	return c.value
}

func (c CollectInstanceStatisticRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CollectInstanceStatisticRequestEngine) UnmarshalJSON(b []byte) error {
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
