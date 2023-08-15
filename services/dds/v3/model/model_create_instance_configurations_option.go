package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateInstanceConfigurationsOption 参数组配置信息。
type CreateInstanceConfigurationsOption struct {

	// 节点类型。 取值：   - 集群实例包含mongos、shard和config节点，各节点下该参数取值分别为“mongos”、“shard”和“config”。   - 副本集实例下该参数取值为“replica”。   - 单节点实例下该参数取值为“single”。
	Type CreateInstanceConfigurationsOptionType `json:"type"`

	// 参数组id。
	ConfigurationId string `json:"configuration_id"`
}

func (o CreateInstanceConfigurationsOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceConfigurationsOption struct{}"
	}

	return strings.Join([]string{"CreateInstanceConfigurationsOption", string(data)}, " ")
}

type CreateInstanceConfigurationsOptionType struct {
	value string
}

type CreateInstanceConfigurationsOptionTypeEnum struct {
	MONGOS  CreateInstanceConfigurationsOptionType
	SHARD   CreateInstanceConfigurationsOptionType
	CONFIG  CreateInstanceConfigurationsOptionType
	REPLICA CreateInstanceConfigurationsOptionType
	SINGLE  CreateInstanceConfigurationsOptionType
}

func GetCreateInstanceConfigurationsOptionTypeEnum() CreateInstanceConfigurationsOptionTypeEnum {
	return CreateInstanceConfigurationsOptionTypeEnum{
		MONGOS: CreateInstanceConfigurationsOptionType{
			value: "mongos",
		},
		SHARD: CreateInstanceConfigurationsOptionType{
			value: "shard",
		},
		CONFIG: CreateInstanceConfigurationsOptionType{
			value: "config",
		},
		REPLICA: CreateInstanceConfigurationsOptionType{
			value: "replica",
		},
		SINGLE: CreateInstanceConfigurationsOptionType{
			value: "single",
		},
	}
}

func (c CreateInstanceConfigurationsOptionType) Value() string {
	return c.value
}

func (c CreateInstanceConfigurationsOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceConfigurationsOptionType) UnmarshalJSON(b []byte) error {
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
