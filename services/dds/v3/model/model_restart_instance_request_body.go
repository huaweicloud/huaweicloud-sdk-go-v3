package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RestartInstanceRequestBody struct {

	// **参数解释：** 待重启对象的类型。 **约束限制：** 重启集群实例下的节点或组时，该参数必选。 - 重启mongos节点时，取值为“mongos”。 - 重启shard组时，取值为“shard”。 - 重启config组时，取值为“config”。 - 重启实例（集群、副本集、单节点）时，不传该参数。 **取值范围：** - mongos - shard - config **默认取值：** 不涉及。
	TargetType *RestartInstanceRequestBodyTargetType `json:"target_type,omitempty"`

	// **参数解释：** 待重启对象的ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。 **约束限制：** 节点状态为正常时，不允许重启主节点。 **取值范围：** - 重启整个实例时，取值为实例ID。 - 重启集群实例shard或config组时取值为shard或config的组ID。 - 重启单个节点时，取值为对应节点的节点ID。 **默认取值：** 不涉及。
	TargetId string `json:"target_id"`

	// **参数解释：** 是否选择节点串行重启。 **约束限制：** 只支持副本集实例。 **取值范围：** - true：表示节点串行重启。 - false：表示节点并行重启。 **默认取值：** false。
	IsSerial *bool `json:"is_serial,omitempty"`

	// **参数解释：** 是否强制重启。 **约束限制：** 仅支持节点状态为异常时进行强制重启。只读节点暂不支持强制重启。 **取值范围：** - true：表示异常节点进行强制重启。 - false：表示进行正常重启。 **默认取值：** false。
	IsForce *bool `json:"is_force,omitempty"`
}

func (o RestartInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"RestartInstanceRequestBody", string(data)}, " ")
}

type RestartInstanceRequestBodyTargetType struct {
	value string
}

type RestartInstanceRequestBodyTargetTypeEnum struct {
	MONGOS RestartInstanceRequestBodyTargetType
	SHARD  RestartInstanceRequestBodyTargetType
	CONFIG RestartInstanceRequestBodyTargetType
}

func GetRestartInstanceRequestBodyTargetTypeEnum() RestartInstanceRequestBodyTargetTypeEnum {
	return RestartInstanceRequestBodyTargetTypeEnum{
		MONGOS: RestartInstanceRequestBodyTargetType{
			value: "mongos",
		},
		SHARD: RestartInstanceRequestBodyTargetType{
			value: "shard",
		},
		CONFIG: RestartInstanceRequestBodyTargetType{
			value: "config",
		},
	}
}

func (c RestartInstanceRequestBodyTargetType) Value() string {
	return c.value
}

func (c RestartInstanceRequestBodyTargetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RestartInstanceRequestBodyTargetType) UnmarshalJSON(b []byte) error {
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
