package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ComponentSnapshotKindObj API类型，固定值“ComponentSnapshot”，该值不可修改。
type ComponentSnapshotKindObj struct {
	value string
}

type ComponentSnapshotKindObjEnum struct {
	COMPONENT_SNAPSHOT ComponentSnapshotKindObj
}

func GetComponentSnapshotKindObjEnum() ComponentSnapshotKindObjEnum {
	return ComponentSnapshotKindObjEnum{
		COMPONENT_SNAPSHOT: ComponentSnapshotKindObj{
			value: "ComponentSnapshot",
		},
	}
}

func (c ComponentSnapshotKindObj) Value() string {
	return c.value
}

func (c ComponentSnapshotKindObj) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ComponentSnapshotKindObj) UnmarshalJSON(b []byte) error {
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
