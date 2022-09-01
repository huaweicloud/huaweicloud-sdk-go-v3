package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 参数组配置信息。
type RestoreNewInstanceConfigurationsOption struct {

	// 节点类型。 取值：   - 集群实例包含mongos、shard和config节点，各节点下该参数取值分别为“mongos”、“shard”和“config”。   - 副本集实例下该参数取值为“replica”。   - 单节点实例下该参数取值为“single”。
	Type RestoreNewInstanceConfigurationsOptionType `json:"type" xml:"type"`

	// 参数组id。
	ConfigurationId string `json:"configuration_id" xml:"configuration_id"`
}

func (o RestoreNewInstanceConfigurationsOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreNewInstanceConfigurationsOption struct{}"
	}

	return strings.Join([]string{"RestoreNewInstanceConfigurationsOption", string(data)}, " ")
}

type RestoreNewInstanceConfigurationsOptionType struct {
	value string
}

type RestoreNewInstanceConfigurationsOptionTypeEnum struct {
	MONGOS  RestoreNewInstanceConfigurationsOptionType
	SHARD   RestoreNewInstanceConfigurationsOptionType
	CONFIG  RestoreNewInstanceConfigurationsOptionType
	REPLICA RestoreNewInstanceConfigurationsOptionType
	SINGLE  RestoreNewInstanceConfigurationsOptionType
}

func GetRestoreNewInstanceConfigurationsOptionTypeEnum() RestoreNewInstanceConfigurationsOptionTypeEnum {
	return RestoreNewInstanceConfigurationsOptionTypeEnum{
		MONGOS: RestoreNewInstanceConfigurationsOptionType{
			value: "mongos",
		},
		SHARD: RestoreNewInstanceConfigurationsOptionType{
			value: "shard",
		},
		CONFIG: RestoreNewInstanceConfigurationsOptionType{
			value: "config",
		},
		REPLICA: RestoreNewInstanceConfigurationsOptionType{
			value: "replica",
		},
		SINGLE: RestoreNewInstanceConfigurationsOptionType{
			value: "single",
		},
	}
}

func (c RestoreNewInstanceConfigurationsOptionType) Value() string {
	return c.value
}

func (c RestoreNewInstanceConfigurationsOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RestoreNewInstanceConfigurationsOptionType) UnmarshalJSON(b []byte) error {
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
