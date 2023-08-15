package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DataEvent struct {
	value string
}

type DataEventEnum struct {
	READ  DataEvent
	WRITE DataEvent
}

func GetDataEventEnum() DataEventEnum {
	return DataEventEnum{
		READ: DataEvent{
			value: "READ",
		},
		WRITE: DataEvent{
			value: "WRITE",
		},
	}
}

func (c DataEvent) Value() string {
	return c.value
}

func (c DataEvent) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataEvent) UnmarshalJSON(b []byte) error {
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
