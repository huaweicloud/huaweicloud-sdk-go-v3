package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
