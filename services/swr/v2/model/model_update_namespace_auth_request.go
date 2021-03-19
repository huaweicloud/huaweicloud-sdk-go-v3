package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type UpdateNamespaceAuthRequest struct {
	ContentType UpdateNamespaceAuthRequestContentType `json:"Content-Type"`

	Namespace string `json:"namespace"`

	Body *[]UserAuth `json:"body,omitempty"`
}

func (o UpdateNamespaceAuthRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateNamespaceAuthRequest struct{}"
	}

	return strings.Join([]string{"UpdateNamespaceAuthRequest", string(data)}, " ")
}

type UpdateNamespaceAuthRequestContentType struct {
	value string
}

type UpdateNamespaceAuthRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 UpdateNamespaceAuthRequestContentType
	APPLICATION_JSON             UpdateNamespaceAuthRequestContentType
}

func GetUpdateNamespaceAuthRequestContentTypeEnum() UpdateNamespaceAuthRequestContentTypeEnum {
	return UpdateNamespaceAuthRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: UpdateNamespaceAuthRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: UpdateNamespaceAuthRequestContentType{
			value: "application/json",
		},
	}
}

func (c UpdateNamespaceAuthRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateNamespaceAuthRequestContentType) UnmarshalJSON(b []byte) error {
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
