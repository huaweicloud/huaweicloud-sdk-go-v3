package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 批量重启边缘实例对象
type BatchReboot struct {
	// 待重启的边缘实例列表。

	Servers *[]BaseId `json:"servers,omitempty"`
	// 重启类型：   - SOFT：普通重启。  - HARD：强制重启。  > 重启必须指定重启类型。

	Type *BatchRebootType `json:"type,omitempty"`
}

func (o BatchReboot) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchReboot struct{}"
	}

	return strings.Join([]string{"BatchReboot", string(data)}, " ")
}

type BatchRebootType struct {
	value string
}

type BatchRebootTypeEnum struct {
	SOFT BatchRebootType
	HARD BatchRebootType
}

func GetBatchRebootTypeEnum() BatchRebootTypeEnum {
	return BatchRebootTypeEnum{
		SOFT: BatchRebootType{
			value: "SOFT",
		},
		HARD: BatchRebootType{
			value: "HARD",
		},
	}
}

func (c BatchRebootType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *BatchRebootType) UnmarshalJSON(b []byte) error {
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
