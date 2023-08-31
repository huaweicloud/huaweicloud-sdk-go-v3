package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEquipmentSpecificConfigRequest Request Object
type ShowEquipmentSpecificConfigRequest struct {

	// 设备类型
	EquipmentType ShowEquipmentSpecificConfigRequestEquipmentType `json:"equipment_type"`
}

func (o ShowEquipmentSpecificConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEquipmentSpecificConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowEquipmentSpecificConfigRequest", string(data)}, " ")
}

type ShowEquipmentSpecificConfigRequestEquipmentType struct {
	value string
}

type ShowEquipmentSpecificConfigRequestEquipmentTypeEnum struct {
	STANDARD ShowEquipmentSpecificConfigRequestEquipmentType
}

func GetShowEquipmentSpecificConfigRequestEquipmentTypeEnum() ShowEquipmentSpecificConfigRequestEquipmentTypeEnum {
	return ShowEquipmentSpecificConfigRequestEquipmentTypeEnum{
		STANDARD: ShowEquipmentSpecificConfigRequestEquipmentType{
			value: "standard",
		},
	}
}

func (c ShowEquipmentSpecificConfigRequestEquipmentType) Value() string {
	return c.value
}

func (c ShowEquipmentSpecificConfigRequestEquipmentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEquipmentSpecificConfigRequestEquipmentType) UnmarshalJSON(b []byte) error {
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
