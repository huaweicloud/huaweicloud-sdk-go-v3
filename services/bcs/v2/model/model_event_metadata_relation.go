package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EventMetadataRelation 事件告警查询条件组合
type EventMetadataRelation struct {

	// 指定查询字段的key，对应metadata里面的key
	Key *string `json:"key,omitempty"`

	// 查询条件中指定key的值
	Value *[]string `json:"value,omitempty"`

	// 该条件与其他条件的组合方式
	Relation *EventMetadataRelationRelation `json:"relation,omitempty"`
}

func (o EventMetadataRelation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventMetadataRelation struct{}"
	}

	return strings.Join([]string{"EventMetadataRelation", string(data)}, " ")
}

type EventMetadataRelationRelation struct {
	value string
}

type EventMetadataRelationRelationEnum struct {
	AND EventMetadataRelationRelation
	OR  EventMetadataRelationRelation
	NOT EventMetadataRelationRelation
}

func GetEventMetadataRelationRelationEnum() EventMetadataRelationRelationEnum {
	return EventMetadataRelationRelationEnum{
		AND: EventMetadataRelationRelation{
			value: "AND",
		},
		OR: EventMetadataRelationRelation{
			value: "OR",
		},
		NOT: EventMetadataRelationRelation{
			value: "NOT",
		},
	}
}

func (c EventMetadataRelationRelation) Value() string {
	return c.value
}

func (c EventMetadataRelationRelation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventMetadataRelationRelation) UnmarshalJSON(b []byte) error {
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
