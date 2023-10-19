package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SortDir desc(降序), asc(升序).
type SortDir struct {
	value string
}

type SortDirEnum struct {
	DESC SortDir
	ASC  SortDir
}

func GetSortDirEnum() SortDirEnum {
	return SortDirEnum{
		DESC: SortDir{
			value: "desc",
		},
		ASC: SortDir{
			value: "asc",
		},
	}
}

func (c SortDir) Value() string {
	return c.value
}

func (c SortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SortDir) UnmarshalJSON(b []byte) error {
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
