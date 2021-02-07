package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type ListApiVersionsRequest struct {
	ContentType ListApiVersionsRequestContentType `json:"Content-Type"`
}

func (o ListApiVersionsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApiVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListApiVersionsRequest", string(data)}, " ")
}

type ListApiVersionsRequestContentType struct {
	value string
}

type ListApiVersionsRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListApiVersionsRequestContentType
	APPLICATION_JSON             ListApiVersionsRequestContentType
}

func GetListApiVersionsRequestContentTypeEnum() ListApiVersionsRequestContentTypeEnum {
	return ListApiVersionsRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListApiVersionsRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListApiVersionsRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListApiVersionsRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListApiVersionsRequestContentType) UnmarshalJSON(b []byte) error {
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
