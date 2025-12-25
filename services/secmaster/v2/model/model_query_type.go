package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QueryType **参数解释**: 查询类型 - SQL SQL查询 - CBSL CBSL查询  **约束限制** 不涉及 **取值范围**: - SQL - CBSL  **默认值** 不涉及
type QueryType struct {
	value string
}

type QueryTypeEnum struct {
	SQL  QueryType
	CBSL QueryType
}

func GetQueryTypeEnum() QueryTypeEnum {
	return QueryTypeEnum{
		SQL: QueryType{
			value: "SQL",
		},
		CBSL: QueryType{
			value: "CBSL",
		},
	}
}

func (c QueryType) Value() string {
	return c.value
}

func (c QueryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryType) UnmarshalJSON(b []byte) error {
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
