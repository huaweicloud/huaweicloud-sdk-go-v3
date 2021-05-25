package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListNamespacesRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType ListNamespacesRequestContentType `json:"Content-Type"`
	// 组织名称

	Namespace *string `json:"namespace,omitempty"`
}

func (o ListNamespacesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListNamespacesRequest struct{}"
	}

	return strings.Join([]string{"ListNamespacesRequest", string(data)}, " ")
}

type ListNamespacesRequestContentType struct {
	value string
}

type ListNamespacesRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListNamespacesRequestContentType
	APPLICATION_JSON             ListNamespacesRequestContentType
}

func GetListNamespacesRequestContentTypeEnum() ListNamespacesRequestContentTypeEnum {
	return ListNamespacesRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListNamespacesRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListNamespacesRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListNamespacesRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListNamespacesRequestContentType) UnmarshalJSON(b []byte) error {
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
