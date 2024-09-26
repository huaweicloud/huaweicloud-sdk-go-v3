package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResizeInstanceOption struct {

	// 对象类型。 - 对于集群实例，该参数为必选。变更mongos节点规格时，取值为“mongos”；变更单个shard组规格、或者批量变更多个shard组规格时，取值为“shard”，变更config组规格时，取值为\"config\"。 - 对于副本集实例，不传该参数。变更readonly节点规格时,取值为“readonly”。 - 对于单节点实例，不传该参数。
	TargetType *ResizeInstanceOptionTargetType `json:"target_type,omitempty"`

	// 待变更规格的节点ID或实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。 - 对于集群实例，变更mongos节点规格时，取值为mongos节点ID；变更单个shard组规格时，取值为shard组ID；批量变更多个shard组规格时，不传该参数；变更config组规格时，取值为config组的ID。 - 对于副本集实例，取值为相应的实例ID。变更readonly节点规格时，取值为readonly节点ID。 - 对于单节点实例，取值为相应的实例ID。
	TargetId string `json:"target_id"`

	// 变更至新规格的资源规格编码。
	TargetSpecCode string `json:"target_spec_code"`
}

func (o ResizeInstanceOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceOption struct{}"
	}

	return strings.Join([]string{"ResizeInstanceOption", string(data)}, " ")
}

type ResizeInstanceOptionTargetType struct {
	value string
}

type ResizeInstanceOptionTargetTypeEnum struct {
	MONGOS   ResizeInstanceOptionTargetType
	SHARD    ResizeInstanceOptionTargetType
	CONFIG   ResizeInstanceOptionTargetType
	READONLY ResizeInstanceOptionTargetType
}

func GetResizeInstanceOptionTargetTypeEnum() ResizeInstanceOptionTargetTypeEnum {
	return ResizeInstanceOptionTargetTypeEnum{
		MONGOS: ResizeInstanceOptionTargetType{
			value: "mongos",
		},
		SHARD: ResizeInstanceOptionTargetType{
			value: "shard",
		},
		CONFIG: ResizeInstanceOptionTargetType{
			value: "config",
		},
		READONLY: ResizeInstanceOptionTargetType{
			value: "readonly",
		},
	}
}

func (c ResizeInstanceOptionTargetType) Value() string {
	return c.value
}

func (c ResizeInstanceOptionTargetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResizeInstanceOptionTargetType) UnmarshalJSON(b []byte) error {
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
