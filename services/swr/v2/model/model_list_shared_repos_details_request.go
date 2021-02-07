package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type ListSharedReposDetailsRequest struct {
	ContentType ListSharedReposDetailsRequestContentType `json:"Content-Type"`
	Filter      *string                                  `json:"filter,omitempty"`
}

func (o ListSharedReposDetailsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSharedReposDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListSharedReposDetailsRequest", string(data)}, " ")
}

type ListSharedReposDetailsRequestContentType struct {
	value string
}

type ListSharedReposDetailsRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListSharedReposDetailsRequestContentType
	APPLICATION_JSON             ListSharedReposDetailsRequestContentType
}

func GetListSharedReposDetailsRequestContentTypeEnum() ListSharedReposDetailsRequestContentTypeEnum {
	return ListSharedReposDetailsRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListSharedReposDetailsRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListSharedReposDetailsRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListSharedReposDetailsRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListSharedReposDetailsRequestContentType) UnmarshalJSON(b []byte) error {
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
