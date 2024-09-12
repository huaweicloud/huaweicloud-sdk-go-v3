package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HitCondition 命中条件配置
type HitCondition struct {

	// **参数解释**： 条件关系。对于多个条件的逻辑运算关系。 **约束限制**： 不涉及。 **取值范围**： * AND：表示多个条件同时满足。 * OR：表示多个条件满足其一即可。 * RESERVED：兜底回复不会去判断其他命中条件。  **默认取值**： 不涉及。
	Relation *HitConditionRelation `json:"relation,omitempty"`

	// **参数解释**： 优先级，数值越低优先级越高；取值0-999，默认值为500，为可选值 **约束限制**： 不涉及
	Priority *int32 `json:"priority,omitempty"`

	// 匹配关系配置
	Tags *[]HitConditionTag `json:"tags,omitempty"`
}

func (o HitCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HitCondition struct{}"
	}

	return strings.Join([]string{"HitCondition", string(data)}, " ")
}

type HitConditionRelation struct {
	value string
}

type HitConditionRelationEnum struct {
	AND      HitConditionRelation
	OR       HitConditionRelation
	RESERVED HitConditionRelation
}

func GetHitConditionRelationEnum() HitConditionRelationEnum {
	return HitConditionRelationEnum{
		AND: HitConditionRelation{
			value: "AND",
		},
		OR: HitConditionRelation{
			value: "OR",
		},
		RESERVED: HitConditionRelation{
			value: "RESERVED",
		},
	}
}

func (c HitConditionRelation) Value() string {
	return c.value
}

func (c HitConditionRelation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HitConditionRelation) UnmarshalJSON(b []byte) error {
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
