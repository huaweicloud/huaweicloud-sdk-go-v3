package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RequestStatus An enumeration. - created - dispatched - completed - timeout - unknown
type RequestStatus struct {
	value string
}

type RequestStatusEnum struct {
	CREATED    RequestStatus
	DISPATCHED RequestStatus
	COMPLETED  RequestStatus
	TIMEOUT    RequestStatus
	UNKNOWN    RequestStatus
}

func GetRequestStatusEnum() RequestStatusEnum {
	return RequestStatusEnum{
		CREATED: RequestStatus{
			value: "created",
		},
		DISPATCHED: RequestStatus{
			value: "dispatched",
		},
		COMPLETED: RequestStatus{
			value: "completed",
		},
		TIMEOUT: RequestStatus{
			value: "timeout",
		},
		UNKNOWN: RequestStatus{
			value: "unknown",
		},
	}
}

func (c RequestStatus) Value() string {
	return c.value
}

func (c RequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RequestStatus) UnmarshalJSON(b []byte) error {
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
