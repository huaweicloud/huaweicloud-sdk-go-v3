package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 实例规格详情。
type RestoreNewInstanceFlavorOption struct {
	// 节点类型。 取值：   - 集群实例包含mongos、shard和config节点，各节点下该参数取值分别为“mongos”、“shard”和“config”。   - 副本集实例下该参数取值为“replica”。   - 单节点实例下该参数取值为“single”。

	Type RestoreNewInstanceFlavorOptionType `json:"type"`
	// 节点数量。 取值：   - 集群实例下“mongos”类型的节点数量可取2~16。   - 集群实例下“shard”类型的组数量可取2~16。   - “shard”类型的组数量可取2~16。   - “config”类型的组数量只能取1。   - “replica”类型的组数量只能取1。   - “single”类型的节点数量只能取1。

	Num string `json:"num"`
	// 磁盘大小。 取值：必须为10的整数倍。单位为GB。   - 对于集群实例，shard组可取10GB~2000GB，config组仅可取20GB。mongos节点不涉及选择磁盘，该参数无意义。   - 对于副本集实例，可取10GB~2000GB。   - 对于单节点实例，可取10GB~1000GB。

	Size *string `json:"size,omitempty"`
	// 资源规格编码

	SpecCode string `json:"spec_code"`
}

func (o RestoreNewInstanceFlavorOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreNewInstanceFlavorOption struct{}"
	}

	return strings.Join([]string{"RestoreNewInstanceFlavorOption", string(data)}, " ")
}

type RestoreNewInstanceFlavorOptionType struct {
	value string
}

type RestoreNewInstanceFlavorOptionTypeEnum struct {
	MONGOS  RestoreNewInstanceFlavorOptionType
	SHARD   RestoreNewInstanceFlavorOptionType
	CONFIG  RestoreNewInstanceFlavorOptionType
	REPLICA RestoreNewInstanceFlavorOptionType
	SINGLE  RestoreNewInstanceFlavorOptionType
}

func GetRestoreNewInstanceFlavorOptionTypeEnum() RestoreNewInstanceFlavorOptionTypeEnum {
	return RestoreNewInstanceFlavorOptionTypeEnum{
		MONGOS: RestoreNewInstanceFlavorOptionType{
			value: "mongos",
		},
		SHARD: RestoreNewInstanceFlavorOptionType{
			value: "shard",
		},
		CONFIG: RestoreNewInstanceFlavorOptionType{
			value: "config",
		},
		REPLICA: RestoreNewInstanceFlavorOptionType{
			value: "replica",
		},
		SINGLE: RestoreNewInstanceFlavorOptionType{
			value: "single",
		},
	}
}

func (c RestoreNewInstanceFlavorOptionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *RestoreNewInstanceFlavorOptionType) UnmarshalJSON(b []byte) error {
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
