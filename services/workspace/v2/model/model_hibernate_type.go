package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HibernateType ECS休眠类型 - SUSPEND: 带外冷休眠 - PAUSE: 带外热休眠
type HibernateType struct {
	value string
}

type HibernateTypeEnum struct {
	SUSPEND HibernateType
	PAUSE   HibernateType
}

func GetHibernateTypeEnum() HibernateTypeEnum {
	return HibernateTypeEnum{
		SUSPEND: HibernateType{
			value: "SUSPEND",
		},
		PAUSE: HibernateType{
			value: "PAUSE",
		},
	}
}

func (c HibernateType) Value() string {
	return c.value
}

func (c HibernateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HibernateType) UnmarshalJSON(b []byte) error {
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
