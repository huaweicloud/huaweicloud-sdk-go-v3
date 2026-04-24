package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreatedBy struct {

	// 工作负载身份创建者类型
	Type CreatedByType `json:"type"`

	// 工作负载身份创建者标识
	Identifier string `json:"identifier"`
}

func (o CreatedBy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatedBy struct{}"
	}

	return strings.Join([]string{"CreatedBy", string(data)}, " ")
}

type CreatedByType struct {
	value string
}

type CreatedByTypeEnum struct {
	CUSTOMER CreatedByType
	SERVICE  CreatedByType
}

func GetCreatedByTypeEnum() CreatedByTypeEnum {
	return CreatedByTypeEnum{
		CUSTOMER: CreatedByType{
			value: "CUSTOMER",
		},
		SERVICE: CreatedByType{
			value: "SERVICE",
		},
	}
}

func (c CreatedByType) Value() string {
	return c.value
}

func (c CreatedByType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatedByType) UnmarshalJSON(b []byte) error {
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
