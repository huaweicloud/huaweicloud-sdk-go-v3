package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListErrorLogsRequest struct {
	InstanceId string `json:"instance_id"`

	StartDate string `json:"start_date"`

	EndDate string `json:"end_date"`

	NodeId *string `json:"node_id,omitempty"`

	Type *ListErrorLogsRequestType `json:"type,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListErrorLogsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListErrorLogsRequest struct{}"
	}

	return strings.Join([]string{"ListErrorLogsRequest", string(data)}, " ")
}

type ListErrorLogsRequestType struct {
	value string
}

type ListErrorLogsRequestTypeEnum struct {
	WARNING ListErrorLogsRequestType
	ERROR   ListErrorLogsRequestType
}

func GetListErrorLogsRequestTypeEnum() ListErrorLogsRequestTypeEnum {
	return ListErrorLogsRequestTypeEnum{
		WARNING: ListErrorLogsRequestType{
			value: "WARNING",
		},
		ERROR: ListErrorLogsRequestType{
			value: "ERROR",
		},
	}
}

func (c ListErrorLogsRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListErrorLogsRequestType) UnmarshalJSON(b []byte) error {
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
