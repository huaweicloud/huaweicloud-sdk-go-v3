package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CreateRepoRequest struct {
	ContentType CreateRepoRequestContentType `json:"Content-Type"`

	Namespace string `json:"namespace"`

	Body *CreateRepoRequestBody `json:"body,omitempty"`
}

func (o CreateRepoRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRepoRequest struct{}"
	}

	return strings.Join([]string{"CreateRepoRequest", string(data)}, " ")
}

type CreateRepoRequestContentType struct {
	value string
}

type CreateRepoRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 CreateRepoRequestContentType
	APPLICATION_JSON             CreateRepoRequestContentType
}

func GetCreateRepoRequestContentTypeEnum() CreateRepoRequestContentTypeEnum {
	return CreateRepoRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: CreateRepoRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: CreateRepoRequestContentType{
			value: "application/json",
		},
	}
}

func (c CreateRepoRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateRepoRequestContentType) UnmarshalJSON(b []byte) error {
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
