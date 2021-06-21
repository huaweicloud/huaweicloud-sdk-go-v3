package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// request
type RefreshTaskRequestRefreshTask struct {
	// 刷新的类型，其值可以为file 或directory，默认为file

	Type *RefreshTaskRequestRefreshTaskType `json:"type,omitempty"`
	// 刷新urls

	Urls []string `json:"urls"`
}

func (o RefreshTaskRequestRefreshTask) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RefreshTaskRequestRefreshTask struct{}"
	}

	return strings.Join([]string{"RefreshTaskRequestRefreshTask", string(data)}, " ")
}

type RefreshTaskRequestRefreshTaskType struct {
	value string
}

type RefreshTaskRequestRefreshTaskTypeEnum struct {
	FILE      RefreshTaskRequestRefreshTaskType
	DIRECTORY RefreshTaskRequestRefreshTaskType
}

func GetRefreshTaskRequestRefreshTaskTypeEnum() RefreshTaskRequestRefreshTaskTypeEnum {
	return RefreshTaskRequestRefreshTaskTypeEnum{
		FILE: RefreshTaskRequestRefreshTaskType{
			value: "file",
		},
		DIRECTORY: RefreshTaskRequestRefreshTaskType{
			value: "directory",
		},
	}
}

func (c RefreshTaskRequestRefreshTaskType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *RefreshTaskRequestRefreshTaskType) UnmarshalJSON(b []byte) error {
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
