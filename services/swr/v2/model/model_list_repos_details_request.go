package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type ListReposDetailsRequest struct {
	ContentType ListReposDetailsRequestContentType `json:"Content-Type"`
	Namespace   *string                            `json:"namespace,omitempty"`
	Name        *string                            `json:"name,omitempty"`
	Category    *string                            `json:"category,omitempty"`
	Filter      *string                            `json:"filter,omitempty"`
}

func (o ListReposDetailsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListReposDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListReposDetailsRequest", string(data)}, " ")
}

type ListReposDetailsRequestContentType struct {
	value string
}

type ListReposDetailsRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListReposDetailsRequestContentType
	APPLICATION_JSON             ListReposDetailsRequestContentType
}

func GetListReposDetailsRequestContentTypeEnum() ListReposDetailsRequestContentTypeEnum {
	return ListReposDetailsRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListReposDetailsRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListReposDetailsRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListReposDetailsRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListReposDetailsRequestContentType) UnmarshalJSON(b []byte) error {
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
