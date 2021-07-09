package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// This is a auto restart Body Object
type RestarInstanceInfo struct {
	// 重启的类型，soft或者hard。 - soft表示软重启（只重启进程）。 - hard表示强制重启（重启虚拟机）。

	Type *RestarInstanceInfoType `json:"type,omitempty"`
}

func (o RestarInstanceInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestarInstanceInfo struct{}"
	}

	return strings.Join([]string{"RestarInstanceInfo", string(data)}, " ")
}

type RestarInstanceInfoType struct {
	value string
}

type RestarInstanceInfoTypeEnum struct {
	SOFT RestarInstanceInfoType
	HARD RestarInstanceInfoType
}

func GetRestarInstanceInfoTypeEnum() RestarInstanceInfoTypeEnum {
	return RestarInstanceInfoTypeEnum{
		SOFT: RestarInstanceInfoType{
			value: "soft",
		},
		HARD: RestarInstanceInfoType{
			value: "hard",
		},
	}
}

func (c RestarInstanceInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *RestarInstanceInfoType) UnmarshalJSON(b []byte) error {
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
